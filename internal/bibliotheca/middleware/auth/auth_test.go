package auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func Test_emailFromTokenClaim_success(t *testing.T) {
	var token *jwt.Token

	want := "user@example.com"
	token = &jwt.Token{
		Claims: jwt.MapClaims(map[string]interface{}{
			"email": want,
		}),
	}

	email, ok := emailFromTokenClaims(token)
	if !ok {
		t.Errorf("unable to extract email from claim")
	}

	if email != "user@example.com" {
		t.Errorf("extracted value(%s) is dirrent from expected value(%s)", email, want)
	}
}

func Test_emailFromTokenClaim_fail(t *testing.T) {
	var token *jwt.Token

	token = &jwt.Token{
		Claims: jwt.MapClaims(map[string]interface{}{})}

	if email, ok := emailFromTokenClaims(token); ok {
		t.Errorf("unexpectedly success to extract value: %s", email)
	}
}
