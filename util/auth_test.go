package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)









/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

FUNCTION_DEF=func ValidateToken(ctx *gin.Context, token string) error 

 */
func TestValidateToken(t *testing.T) {

	testCases := []struct {
		name            string
		token           string
		prepareToken    bool
		expectedErr     bool
		expectedErrType error
	}{
		{
			name:            "Valid token",
			token:           "valid_token",
			prepareToken:    true,
			expectedErr:     false,
			expectedErrType: nil,
		},
		{
			name:            "Expired token",
			token:           "expired_token",
			prepareToken:    true,
			expectedErr:     true,
			expectedErrType: jwt.ErrSignatureInvalid,
		},
		{
			name:            "Invalid signature",
			token:           "invalid_signature",
			prepareToken:    true,
			expectedErr:     true,
			expectedErrType: jwt.ErrSignatureInvalid,
		},
		{
			name:            "No token",
			token:           "",
			prepareToken:    false,
			expectedErr:     true,
			expectedErrType: errors.New("Token not found"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := getMockedContext()
			if tt.prepareToken {
				c.Request.Header.Set("Authorization", tt.token)
			}

			actualErr := ValidateToken(c, tt.token)

			if tt.expectedErr == false && actualErr != nil {
				t.Errorf("Expected nil, but got error: %v", actualErr)
			}

			if tt.expectedErr == true {
				if actualErr != nil {
					if actualErr.Error() != tt.expectedErrType.Error() {
						t.Errorf("Expected error type: %v, but got:  %v", tt.expectedErrType, actualErr)
					}

				} else {
					t.Errorf("Expected error: %v, but got nil", tt.expectedErrType)
				}
			}
		})
	}
}

func getMockedContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	return c, w
}

