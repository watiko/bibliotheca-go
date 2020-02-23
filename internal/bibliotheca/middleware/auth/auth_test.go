package auth_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
)

func TestEmailFromTokenClaim(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		emailFromTokenClaimSuccess(t)
	})
	t.Run("Fail", func(t *testing.T) {
		emailFromTokenClaimFail(t)
	})
}

func emailFromTokenClaimSuccess(t *testing.T) {
	t.Helper()
	var token *jwt.Token

	want := "user@example.com"
	token = &jwt.Token{
		Claims: jwt.MapClaims(map[string]interface{}{
			"email": want,
		}),
	}

	email, ok := ExportEmailFromTokenClaims(token)
	assert.True(t, ok, "unable to extract email from claim")
	assert.Equal(t, want, email)
}

func emailFromTokenClaimFail(t *testing.T) {
	t.Helper()
	var token *jwt.Token

	token = &jwt.Token{
		Claims: jwt.MapClaims(map[string]interface{}{}),
	}

	email, ok := ExportEmailFromTokenClaims(token)
	assert.Falsef(t, ok, "want nil, but got %s", email)
}

func credentialMockedAuth(t *testing.T) (*httptest.Server, gin.HandlerFunc) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, firebaseMock.authCredentials)
	}))

	f := ExportNewFirebaseKeyGetter(ts.URL)
	v := ExportNewValidationKeyGetter(f)
	mw := ExportNewJWTMiddlewareWithValidationKeyGetter(v)
	auth := ExportAuthWithJWTMiddleware(mw)

	return ts, auth
}

func TestAuthMiddleware(t *testing.T) {
	ts, auth := credentialMockedAuth()
	defer ts.Close()

	router := gin.New()
	router.Use(auth)

	router.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	t.Run("valid JWT", func(t *testing.T) {
		defer firebaseMock.resetJWTTime()
		firebaseMock.setValidTime()

		w := performRequest(router, "GET", "/ok", firebaseMock.jwtHeader())
		assert.Equal(t, "ok", w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("expired JWT", func(t *testing.T) {
		defer firebaseMock.resetJWTTime()
		firebaseMock.setExpiredTime()

		w := performRequest(router, "GET", "/ok", firebaseMock.jwtHeader())
		assert.Equal(t, "Token is expired\n", w.Body.String())
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("without token", func(t *testing.T) {
		defer firebaseMock.resetJWTTime()
		firebaseMock.setExpiredTime()

		w := performRequest(router, "GET", "/ok")
		assert.Equal(t, "Required authorization token not found\n", w.Body.String())
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("invalid token", func(t *testing.T) {
		defer firebaseMock.resetJWTTime()
		firebaseMock.setValidTime()

		w := performRequest(router, "GET", "/ok", header{Key: "Authorization", Value: "Bearer invalid"})
		assert.Equal(t, "token contains an invalid number of segments\n", w.Body.String())
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("user context", func(t *testing.T) {
		defer firebaseMock.resetJWTTime()
		firebaseMock.setValidTime()

		var user *User
		router.GET("/set_user_context", func(c *gin.Context) {
			user, _ = GetUser(c)
			c.Status(http.StatusOK)
		})

		w := performRequest(router, "GET", "/set_user_context", firebaseMock.jwtHeader())

		assert.Equal(t, "", w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
		if ok := assert.NotNil(t, user); !ok {
			return
		}
		assert.Equal(t, "watiko@mail.watiko.net", user.Email)
	})
}

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type FirebaseMock struct {
	jwt                string
	authCredentials    string
	jwtIssuedTimestamp int64
}

var firebaseMock = FirebaseMock{
	// expired
	jwt: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImNiOGUwZDk3Mjg2MWIwNGJlN2RjNzVhMWIzYmUzYjIyOWIyNWYyMDUiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoid2F0aWtvIiwicGljdHVyZSI6Imh0dHBzOi8vbGg2Lmdvb2dsZXVzZXJjb250ZW50LmNvbS8tSV9UempUZTl4cEEvQUFBQUFBQUFBQUkvQUFBQUFBQUFBUUkva1VXbWY5UlIzNVUvcGhvdG8uanBnIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2JpYmxpb3RoZWNhLXdhdGlrbyIsImF1ZCI6ImJpYmxpb3RoZWNhLXdhdGlrbyIsImF1dGhfdGltZSI6MTU2NTYyNTg4MSwidXNlcl9pZCI6IjB6TUNUYk85Q3RkYnVaWGtxS1lQeEU0ZmdRWTIiLCJzdWIiOiIwek1DVGJPOUN0ZGJ1WlhrcUtZUHhFNGZnUVkyIiwiaWF0IjoxNTgyNDI3MzE1LCJleHAiOjE1ODI0MzA5MTUsImVtYWlsIjoid2F0aWtvQG1haWwud2F0aWtvLm5ldCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7Imdvb2dsZS5jb20iOlsiMTA1NDE3Mjg3NDY4NTM4NDkxODAzIl0sImVtYWlsIjpbIndhdGlrb0BtYWlsLndhdGlrby5uZXQiXX0sInNpZ25faW5fcHJvdmlkZXIiOiJnb29nbGUuY29tIn19.HZ1K06LEWlrdCsHhoBnrBiuhqMjAnWYI2mdpSMl6KcxS5caWjHn7nTcHhd6_q1CjntndqKxOXpr8ploCElZHRaAEcQr4cUuCBQXkhgNHIDgWiHUhBYl1yzB7lil3QQNkRw3o_SCLv6Hp3aRe4AaF9RczeCpj4aRTPUVd0lzFJq4PBMBIxzMqIYG-3q4HeNlO_8y__lDpAEyPIlKFL3MaiFsfPYbj9xHA4PIZhLrlHR11wBWVeBeVMnWsPxXytS9EaTgG7KYxsVZx705IwDbBsb77nrtKNIuFDLp42nrInW7GYGIGV0DkDoB0N7lIh-1NgBliE_4rOuL1hSjfs9XJMw",
	// https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com
	authCredentials: `{
"cb8e0d972861b04be7dc75a1b3be3b229b25f205": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIIJNWhnK6WHLowDQYJKoZIhvcNAQEFBQAwMTEvMC0GA1UE\nAxMmc2VjdXJldG9rZW4uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wHhcNMjAw\nMjE1MDkxOTU0WhcNMjAwMzAyMjEzNDU0WjAxMS8wLQYDVQQDEyZzZWN1cmV0b2tl\nbi5zeXN0ZW0uZ3NlcnZpY2VhY2NvdW50LmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAJJWKXB6ytAd3sk+vSTwvHJfgj4yAaog79eKGr/TNOAg75rK\nztLpDFvS16rEFlr5kSqENiNc0cVZSFn+avK+YI4NM7slwJhT75aFadVW/pHGG9VA\n/g8L+iOKEM+H6uQ4un+M2k6tDfzWcXgUl4mJhGDAA1/vMID4I1WEueNApdP35AT2\nCiqr5rpRDYKwiStfCVO+SwCNBqd5nuyupvGIBKUZwczjzckcUuSG898dPrOWX2JY\nuuGVJANTdKQLi6505EzzPbuxZgrhBiX7hfjME5yF6ARJHfaHDRHfGcX8rKpDv1Kg\nayP9zISuS5qmrl/BuI/dtoEvfZ6zn5inEZ+Htm8CAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAOBgNVHQ8BAf8EBAMCB4AwFgYDVR0lAQH/BAwwCgYIKwYBBQUHAwIwDQYJ\nKoZIhvcNAQEFBQADggEBAGufBdaKhqDQf3WqyK/6nFUdz+dVI64gbhE1WrcYkKaW\ntXgBWc5nH73nLGGdijaDLtInqnn5z3TEziDE3oaruf6BB78QaTBjCQ4wYKsooija\nytiyP4VujQtEYi/hMFI0+Og9hP2PhufhG8ZMEfus3j+M4IXOuIdCydiQARNgldy4\ntWwoS9TxLzM6QnkZ8ktnTXimZyHCVtJvMZvFC+XWK7b9HNpIsoeJ8j6FrYg+QIKL\nrK8p1LFpeVdi9YtPoinrljQUfX5b6JKjlcLHUCGcKZkPrUnlAeAgUpWOrbnhO6FN\nmG/2oOrbOQvIrhsghIwxkPzqMCu0KTORFZX2ICjKwms=\n-----END CERTIFICATE-----\n",
"0ea3f7a0248be4e0df02aeeeb20b1d2e2b7f2474": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIIbKyKmVqvyMkwDQYJKoZIhvcNAQEFBQAwMTEvMC0GA1UE\nAxMmc2VjdXJldG9rZW4uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wHhcNMjAw\nMjIzMDkxOTU0WhcNMjAwMzEwMjEzNDU0WjAxMS8wLQYDVQQDEyZzZWN1cmV0b2tl\nbi5zeXN0ZW0uZ3NlcnZpY2VhY2NvdW50LmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMdMjAzXgXUqwgJo8S5V6NpuB3gGIxT2zvBUuJGKVLk6E8Nz\nt9INbKL5Bv6eW9tjJqRAzgiC6RL8XsZxpvKTtcTZ9C3V3/+ItbBtJIh8KAkKtLAF\nu2n1YzO6xXtYr44aYw6zArG87iQb+qFER3DH+xx8Nb0jNnCHnrS4/rFR6ZqY+X/v\nT6EgliW27WH3VRlLNI+bqPNPBAmsVRd6FgCwq06al9VCkoPXEgDxnsSjRRrHW93n\nXKDeYPZ7TK4Ayr35D6QcLbUDIm73qOR4gqM1REu02ezG/nE9P0kuG0650iXHT9LT\nLo9I1aiDDof8PBl+EnUg8LRlTIQJzV8zBdMzbEUCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAOBgNVHQ8BAf8EBAMCB4AwFgYDVR0lAQH/BAwwCgYIKwYBBQUHAwIwDQYJ\nKoZIhvcNAQEFBQADggEBAJZIw3F/OsuXel6Acj/7KvXmpXgdP8Rt/sb0O6419Oyi\n+SSPViW6CzoczLeqlQIvUu7pyU9F1bVP/HBQsmk57AntP9dqsyjAUx/ZD9ODFZBY\nlPozLPX7+2aSmjBp8/uSQ1hPFROvC4Td1z26H/b3mze1L1xuXPj32JWsX3hHGh+7\nbP1cjIgllWsGBAczenLBXgEBDN2eS/ESV8UIC7w0QbIYfzcCTdzmwAijkNpHJfcA\nnNfrhca7hrADsBR+t61DYax/ZUtxtBBkQ6HQJyrACvS1shNCayuoSp9z3D9LM8E8\nbiBcIPaIVooX/GjU3PncthjmkxfgWgUp1vio8uNYfa0=\n-----END CERTIFICATE-----\n"
}`,
	jwtIssuedTimestamp: 1582427315,
}

func (f *FirebaseMock) jwtHeader() header {
	return header{
		Key:   "Authorization",
		Value: fmt.Sprintf("Bearer %s", f.jwt),
	}
}

func (f *FirebaseMock) setJWTTime(addSec int64) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(f.jwtIssuedTimestamp+addSec, 0)
	}
}

func (f *FirebaseMock) setValidTime() {
	f.setJWTTime(60)
}

func (f *FirebaseMock) setExpiredTime() {
	f.setJWTTime(60*60 + 1) // Firebase Auth Token expires 1 hour later
}

func (f *FirebaseMock) resetJWTTime() {
	jwt.TimeFunc = time.Now
}
