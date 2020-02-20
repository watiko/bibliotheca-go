package auth

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
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

	email, ok := emailFromTokenClaims(token)
	assert.True(t, ok, "unable to extract email from claim")
	assert.Equal(t, want, email)
}

func emailFromTokenClaimFail(t *testing.T) {
	t.Helper()
	var token *jwt.Token

	token = &jwt.Token{
		Claims: jwt.MapClaims(map[string]interface{}{}),
	}

	email, ok := emailFromTokenClaims(token)
	assert.Truef(t, ok, "want nil, but got %s", email)
}
