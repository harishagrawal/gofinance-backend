package util

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

type validateTokenFunc func(ctx *gin.Context, token string) error

var validateToken validateTokenFunc = ValidateToken

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
		{
			name:  "Authorization Header Not Present",
			token: "",
			mockValidateTokenFn: func(ctx *gin.Context, token string) error {
				return errors.New("Authorization Header is not present")
			},
			expectedError: errors.New("Authorization Header is not present"),
		},
		{
			name:  "Authorization Header Contains Insufficient Fields",
			token: "Bearer",
			mockValidateTokenFn: func(ctx *gin.Context, token string) error {
				return errors.New("Authorization Header is not in correct format")
			},
			expectedError: errors.New("Authorization Header is not in correct format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			validateToken = tt.mockValidateTokenFn

			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.token)
			rec := httptest.NewRecorder()

			r.GET("/", func(c *gin.Context) {
				err := GetTokenInHeaderAndVerify(c)
				assert.Equal(t, tt.expectedError, err)
			})
			r.ServeHTTP(rec, req)
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
		return &jwt.Token{Valid: claims.(*Claims).Username != "invalid"}, nil
	}
	defer func() { jwt.ParseWithClaims = parseWithClaimsOrig }()

	testCases := []struct {
		name        string
		token       string
		mockClaims  *Claims
		expectError bool
		statusCode  int
	}{
		{
			name:        "Token is valid and parsable",
			token:       "valid_token",
			mockClaims:  &Claims{Username: "valid", RegisteredClaims: jwt.RegisteredClaims{}},
			expectError: false,
			statusCode:  http.StatusOK,
		},
		{
			name:        "Token is invalid",
			token:       "invalid_token",
			mockClaims:  &Claims{Username: "invalid", RegisteredClaims: jwt.RegisteredClaims{}},
			expectError: true,
			statusCode:  http.StatusBadRequest,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			err := ValidateToken(ctx, tt.token)

			assert.Equal(t, tt.expectError, err != nil)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}
