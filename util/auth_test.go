package util

import (
	"errors"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"github.com/stretchr/testify/assert"
	"util"
)


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	var tests = []struct {
		name        string
		token       string
		expectedErr error
	}{
		{name: "Valid token in header", token: "Bearer token", expectedErr: nil},
		{name: "Invalid token format in header", token: "token", expectedErr: errors.New("index out of range")},
		{name: "No auth header", token: "", expectedErr: errors.New("index out of range")},
		{name: "Invalid token in header", token: "Bearer inval!dtoken", expectedErr: jwt.ErrSignatureInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request.Header.Set("Authorization", tt.token)

			err := GetTokenInHeaderAndVerify(ctx)

			if err != nil {
				t.Log(err.Error())
				if tt.expectedErr == nil {
					t.Errorf("got an error when none was expected: %s", err.Error())
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("got %s, want %s", err.Error(), tt.expectedErr.Error())
				}
			} else if tt.expectedErr != nil {
				t.Errorf("expected %s, got nil", tt.expectedErr.Error())
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	secretKey := "secret_key"

	validClaims := &Claims{
		Username: "test",
		StandardClaims: jwt.StandardClaims{
			Issuer: "test",
		},
	}

	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, validClaims)
	validTokenString, _ := validToken.SignedString([]byte(secretKey))

	testCases := []struct {
		name           string
		token          string
		expectedError  error
		expectedStatus int
	}{
		{
			name:           "Scenario 1: Validate Valid JWT Token",
			token:          validTokenString,
			expectedError:  nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Scenario 2: Validate Invalid JWT Token",
			token:          strings.ReplaceAll(validTokenString, "7", "8"),
			expectedError:  jwt.ErrSignatureInvalid,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Scenario 3: Validate an Empty JWT Token",
			token:          "",
			expectedError:  jwt.ErrSignatureInvalid,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest("GET", "/", nil)
			request.Header.Set("Authorization", tc.token)
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = request

			err := util.ValidateToken(ctx, tc.token)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedStatus, ctx.Writer.Status())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedStatus, ctx.Writer.Status())
			}
		})
	}
}

