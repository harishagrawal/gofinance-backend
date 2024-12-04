package util

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	var jwtSignedKey = []byte("secret_key")

	tests := []struct {
		name               string
		token              string
		expectedHTTPStatus int
		expectError        bool
	}{
		{
			name:               "ValidateToken function with valid JWT token",
			token:              generateTestToken(),
			expectedHTTPStatus: http.StatusOK,
			expectError:        false,
		},
		{
			name:               "ValidateToken function with invalid JWT token",
			token:              "invalidToken",
			expectedHTTPStatus: http.StatusUnauthorized,
			expectError:        true,
		},
		{
			name:               "ValidateToken function with empty JWT token",
			token:              "",
			expectedHTTPStatus: http.StatusBadRequest,
			expectError:        true,
		},
		{
			name:               "ValidateToken function where JWT signature is invalid",
			token:              generateTestToken()[1:],
			expectedHTTPStatus: http.StatusUnauthorized,
			expectError:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := setupContext(tt.token)

			err := ValidateToken(ctx, tt.token)

			if tt.expectError {
				if assert.Error(t, err) {
					assert.Equal(t, tt.expectedHTTPStatus, ctx.Writer.Status())
				}

			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedHTTPStatus, ctx.Writer.Status())
			}
		})
	}
}

func generateTestToken() string {

	claims := &jwt.StandardClaims{
		Issuer: "testIssuer",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validTestToken, _ := token.SignedString([]byte("secret_key"))
	return validTestToken
}

func setupContext(token string) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	return c
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name           string
		token          string
		isErrExpected  bool
		mockValidation func(token string) error
	}{
		{
			name:          "Valid Token Passed",
			token:         "Bearer validToken",
			isErrExpected: false,
			mockValidation: func(token string) error {
				return nil
			},
		},
		{
			name:          "Invalid Token Passed",
			token:         "Bearer invalidToken",
			isErrExpected: true,
			mockValidation: func(token string) error {
				return fmt.Errorf("invalid token")
			},
		},
		{
			name:          "Malformed Authentication Header",
			token:         "MalformedToken",
			isErrExpected: true,
			mockValidation: func(token string) error {
				return fmt.Errorf("invalid token")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("authorization", tc.token)
			ValidateTokenFunc = tc.mockValidation
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request = req

			err := GetTokenInHeaderAndVerify(ctx)

			if !tc.isErrExpected {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}

	ValidateTokenFunc = ValidateToken
}

