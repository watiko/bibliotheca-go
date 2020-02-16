package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
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

		token := c.Request.Context().Value(userContextKey)

		if token, ok := token.(*jwt.Token); !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			if ok := setUser(c, token); !ok {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}

		c.Next()
	}
}

var userContextKey = "user"

type User struct {
	Email string
	Token *jwt.Token
}

func setUser(c *gin.Context, token *jwt.Token) (ok bool) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if email, ok := claims["email"]; ok {
			if email, ok := email.(string); ok {
				c.Set(userContextKey, &User{
					Email: email,
					Token: token,
				})
				return true
			}
		}
	}
	return false
}

func GetUser(c *gin.Context) (user *User, exists bool) {
	u, exists := c.Get(userContextKey)
	if !exists {
		return nil, exists
	}

	return u.(*User), exists
}

func getKidIfFirebaseAuth(token *jwt.Token) (kid string, ok bool) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if _, ok := claims["firebase"]; ok {
			if kid, ok := token.Header["kid"]; ok {
				if kid, ok := kid.(string); ok {
					return kid, true
				}
			}
		}
	}
	return "", false
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
	if kid, ok := getKidIfFirebaseAuth(token); ok {
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

	return nil, errors.New("invalid JWT Key is passed")
}

func mkJwtMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: validationKeyGetter,
		UserProperty:        userContextKey,
		SigningMethod:       jwt.SigningMethodRS256,
	})
}
