package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/lru"
	"github.com/gregjones/httpcache"
)

func Auth() gin.HandlerFunc {
	return authWithJWTMiddleware(defaultJWTMiddleware())
}

func authWithJWTMiddleware(jwtMiddleware *jwtmiddleware.JWTMiddleware) gin.HandlerFunc {
	if jwtMiddleware == nil {
		jwtMiddleware = defaultJWTMiddleware()
	}

	return func(c *gin.Context) {
		if err := jwtMiddleware.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Next()
			return
		}

		token, ok := c.Request.Context().Value(userContextKey).(*jwt.Token)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			c.Next()
			return
		}

		if ok := setUser(c, token); !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			c.Next()
			return
		}
	}
}

var userContextKey = "user"

type User struct {
	Email string
	Token *jwt.Token
}

func setUser(c *gin.Context, token *jwt.Token) bool {
	email, ok := emailFromTokenClaims(token)
	if !ok {
		return false
	}

	c.Set(userContextKey, &User{
		Email: email,
		Token: token,
	})
	return true
}

func GetUser(c *gin.Context) (*User, bool) {
	_user, exists := c.Get(userContextKey)
	if !exists {
		return nil, exists
	}

	user, ok := _user.(*User)
	return user, ok
}

func emailFromTokenClaims(token *jwt.Token) (string, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	_email, ok := claims["email"]
	if !ok {
		return "", false
	}

	email, ok := _email.(string)
	return email, ok
}

func kidFromTokenHeader(token *jwt.Token) (string, bool) {
	_kid, ok := token.Header["kid"]
	if !ok {
		return "", false
	}

	kid, ok := _kid.(string)
	return kid, ok
}

func isFirebaseAuth(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	if _, ok := claims["firebase"]; !ok {
		return false
	}

	return true
}

func getKidIfFirebaseAuth(token *jwt.Token) (string, bool) {
	if !isFirebaseAuth(token) {
		return "", false
	}

	kid, ok := kidFromTokenHeader(token)
	return kid, ok
}

type firebaseKeyGetter struct {
	firebaseAuthCredentialURL string
	keyCache                  *lru.Cache
	httpCacheTransport        http.RoundTripper
}

var defaultFirebaseAuthCredentialURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

func defaultFirebaseKeyGetter() *firebaseKeyGetter {
	return newFirebaseKeyGetter(defaultFirebaseAuthCredentialURL)
}

func newFirebaseKeyGetter(firebaseAuthCredentialURL string) *firebaseKeyGetter {
	return &firebaseKeyGetter{
		firebaseAuthCredentialURL: firebaseAuthCredentialURL,
		keyCache:                  lru.New(5),
		httpCacheTransport:        httpcache.NewMemoryCacheTransport(),
	}
}

func (f *firebaseKeyGetter) Get(kid string) (interface{}, error) {
	if kid == "" {
		return nil, errors.New("empty kid is passed")
	}

	if key, ok := f.keyCache.Get(kid); ok {
		return key, nil
	}

	creds, err := f.fetchFirebaseAuthCredential()
	if err != nil {
		return nil, err
	}

	cred, ok := (*creds)[kid]
	if !ok {
		return nil, errors.New("JWT key does not found")
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cred))
	if err != nil {
		return nil, err
	}

	f.keyCache.Add(kid, key)
	return key, nil
}

func (f *firebaseKeyGetter) fetchFirebaseAuthCredential() (*map[string]string, error) {
	client := http.Client{Transport: f.httpCacheTransport}
	resp, err := client.Get(f.firebaseAuthCredentialURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var out map[string]string
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

type validationKeyGetterFn = func(*jwt.Token) (interface{}, error)

func newValidationKeyGetter(firebaseKeyGetter *firebaseKeyGetter) validationKeyGetterFn {
	return func(token *jwt.Token) (interface{}, error) {
		kid, ok := getKidIfFirebaseAuth(token)
		if !ok {
			return nil, errors.New("invalid JWT is passed")
		}

		key, err := firebaseKeyGetter.Get(kid)
		return key, err
	}
}

func newJWTMiddlewareWithValidationKeyGetter(keyGetter validationKeyGetterFn) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: keyGetter,
		UserProperty:        userContextKey,
		SigningMethod:       jwt.SigningMethodRS256,
	})
}

func defaultJWTMiddleware() *jwtmiddleware.JWTMiddleware {
	firebaseKeyGetter := defaultFirebaseKeyGetter()
	keyGetter := newValidationKeyGetter(firebaseKeyGetter)
	return newJWTMiddlewareWithValidationKeyGetter(keyGetter)
}
