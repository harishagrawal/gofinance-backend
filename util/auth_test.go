package util

import (
	"errors"
	"net/http"
	"testing"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

type validateTokenFunc func(ctx *gin.Context, token string) error

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	r := gin.Default()

	invalidToken := "invalidToken"
	validToken := "Bearer XYZ123"

	tests := []struct {
		name                string
		token               string
		mockValidateTokenFn validateTokenFunc
		expectedError       error
	}{
	    .... // The rest of your code here
	
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

*/
func TestValidateToken(t *testing.T) {

	parseWithClaimsOrig := jwt.ParseWithClaims
	jwt.ParseWithClaims = func(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
		return &jwt.Token{Valid: claims.(*Claims).Username != "invalid"}, nil
	}
	defer func() { jwt.ParseWithClaims = parseWithClaimsOrig }()
	
	....  // The rest of your code here
}
