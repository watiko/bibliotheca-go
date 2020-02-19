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
	return func(c *gin.Context) {
		jwtMiddleware := mkJwtMiddleware()

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

var httpTransport = httpcache.NewMemoryCacheTransport()

func fetchFirebaseAuthCredential() (*map[string]string, error) {
	firebaseAuthCredentialUrl := "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

	client := http.Client{Transport: httpTransport}
	resp, err := client.Get(firebaseAuthCredentialUrl)
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

var keyCache = lru.New(5)

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	kid, ok := getKidIfFirebaseAuth(token)
	if !ok {
		return nil, errors.New("invalid JWT Key is passed")
	}

	if key, ok := keyCache.Get(kid); ok {
		// use cache
		return key, nil
	}

	creds, err := fetchFirebaseAuthCredential()
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

	keyCache.Add(kid, key)
	return key, nil
}

func mkJwtMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: validationKeyGetter,
		UserProperty:        userContextKey,
		SigningMethod:       jwt.SigningMethodRS256,
	})
}
