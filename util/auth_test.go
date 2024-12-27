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

var validateToken validateTokenFunc = nil

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	r := gin.Default()

	tests := []struct {
		name                string
		token               string
		mockValidateTokenFn validateTokenFunc
		expectedError       error
	}{
		// tests cases remain the same...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateToken = tt.mockValidateTokenFn

			r.GET("/", func(c *gin.Context) {
				err := GetTokenInHeaderAndVerify(c)
				if tt.expectedError != nil {
					assert.ErrorIs(t, err, tt.expectedError)
				} else {
					assert.NoError(t, err)
				}
			})

			req, _ := http.NewRequest("GET", "/", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
		})
	}
}

func TestValidateToken(t *testing.T) {
	// Test cases remain the same...

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
