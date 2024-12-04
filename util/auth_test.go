package util

import (
	"errors"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	mockCtrl, _ := gin.CreateTestContext(nil)

	var jwtSignedKey = []byte("secret_key")
	var claims = &Claims{
		Username: "testUser",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSignedKey)

	scenarios := []struct {
		token     string
		wantError bool
		errorType error
	}{
		{
			token:     tokenString,
			wantError: false,
			errorType: nil,
		},
		{
			token:     "invalidToken123",
			wantError: true,
			errorType: jwt.ErrSignatureInvalid,
		},
		{
			token:     "malformedToken::",
			wantError: true,
			errorType: jwt.ErrSignatureInvalid,
		},
		{

			token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RVc2VyIiwiZXhwIjoiMTYzMzcwMDg0NSJ9.C3eAaUAC2sM8_wWSMztlgjQ_yCzTcyQL4toxzI-xrTI",
			wantError: true,
			errorType: errors.New("Token is expired"),
		},
	}

	for _, tp := range scenarios {
		t.Run("testing validatetoken method", func(t *testing.T) {
			err := ValidateToken(mockCtrl, tp.token)
			if tp.wantError {
				assert.Error(t, err)
				assert.Equal(t, tp.errorType, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name             string
		authorization    string
		setupMockContext func() *gin.Context
		expectedError    error
	}{
		{
			name:          "Valid token",
			authorization: "Bearer eyJ0eXAiOiJqb2tpIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			setupMockContext: func() *gin.Context {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Set("Authorization", "Bearer "+"eyJ0eXAiOiJqb2tpIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
				return &gin.Context{Request: req}
			},
			expectedError: nil,
		},
		{
			name:          "Invalid token",
			authorization: "Bearer invalidtoken",
			setupMockContext: func() *gin.Context {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Set("Authorization", "Bearer "+"invalidtoken")
				return &gin.Context{Request: req}
			},
			expectedError: jwt.ErrSignatureInvalid,
		},
		{
			name:          "Absent authorization header",
			authorization: "",
			setupMockContext: func() *gin.Context {
				req, _ := http.NewRequest("GET", "/", nil)
				return &gin.Context{Request: req}
			},
			expectedError: errors.New("missing authorization header"),
		},
		{
			name:          "Present authorization header but no token",
			authorization: "Bearer ",
			setupMockContext: func() *gin.Context {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Set("Authorization", "Bearer "+"")
				return &gin.Context{Request: req}
			},
			expectedError: jwt.ErrSignatureInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setupMockContext()
			err := GetTokenInHeaderAndVerify(ctx)

			if err != nil && tt.expectedError == nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error: %v, but got nil", tt.expectedError)
				return
			}

			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error: %v, but got error: %v", tt.expectedError, err)
				return
			}
		})
	}
}

