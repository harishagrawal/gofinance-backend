package util

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	"github.com/golang-jwt/jwt/v4"
)

type validateTokenFunc func(ctx *gin.Context, token string) error

var validateTokenFuncMatcher validateTokenFunc = func(ctx *gin.Context, token string) error {
	return ValidateToken(ctx, token)
}

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
		{
			name:  "Successful Token Retrieval and Verification",
			token: validToken,
			mockValidateTokenFn: func(ctx *gin.Context, token string) error {
				return nil
			},
			expectedError: nil,
		},
		{
			name:  "Token Verification Failed",
			token: invalidToken,
			mockValidateTokenFn: func(ctx *gin.Context, token string) error {
				return errors.New("Token is invalid")
			},
			expectedError: errors.New("Token is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			validateTokenFuncMatcher = tt.mockValidateTokenFn

			r.GET("/", func(c *gin.Context) {
				err := GetTokenInHeaderAndVerify(c)
				assert.IsType(t, tt.expectedError, err)
			})

			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.token)

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02
*/
func TestValidateToken(t *testing.T) {

	parseWithClaimsOrig := jwt.ParseWithClaims
	jwt.ParseWithClaims = func(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
		return &jwt.Token{Valid: true}, nil
	}
	defer func() { jwt.ParseWithClaims = parseWithClaimsOrig }()

	invalidToken := "invalid_token"
	validToken := "valid_token"

	testCases := []struct {
		name        string
		token       string
		expectError bool
		statusCode  int
	}{
		{
			name:        "Token is valid and parsable",
			token:       validToken,
			expectError: false,
			statusCode:  http.StatusOK,
		},
		{
			name:        "Token is invalid",
			token:       invalidToken,
			expectError: true,
			statusCode:  http.StatusUnauthorized,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ginctx, _ := gin.CreateTestContext(w)

			err := ValidateToken(ginctx, tt.token)

			assert.Equal(t, tt.expectError, err != nil)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}
