package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"time"
)


var jwtSignedKey = []byte("secret_key")

var mockContext = &gin.Context{}/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	dummyToken, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	).SignedString(jwtSignedKey)

	testCases := []struct {
		name           string
		token          string
		expectedErr    string
		expectedStatus int
	}{
		{
			name:           "successful_authentication",
			token:          dummyToken,
			expectedErr:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid_token_in_header",
			token:          "invalid_token",
			expectedErr:    "Token is invalid",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "missing_authorization_header_key",
			token:          "",
			expectedErr:    "Token is invalid",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "valid_token_with_extra_white_spaces",

			token:          " " + dummyToken + "  ",
			expectedErr:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "empty_authorization_header",
			token:          "",
			expectedErr:    "Token is invalid",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			response := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(response)

			context.Request, _ = http.NewRequest("POST", "/", nil)
			context.Request.Header.Set("Authorization", "Bearer "+tc.token)

			err := GetTokenInHeaderAndVerify(context)

			if tc.expectedErr == "" {
				assert.NoError(t, err, "The test case was expected to return no error")
				assert.Equal(t, tc.expectedStatus, response.Code, "The http response codes do not match")
			} else {
				assert.Error(t, err, "The test cases was expected to return an error")
				assert.Equal(t, tc.expectedErr, err.Error(), "The error messages do not match")
				assert.Equal(t, tc.expectedStatus, response.Code, "The http response codes do not match")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	tests := []struct {
		name          string
		token         string
		expectedError error
	}{
		{
			name:          "Test with valid token",
			token:         makeToken(time.Now().Add(time.Hour * 24)),
			expectedError: nil,
		},
		{
			name:          "Test with invalid token",
			token:         makeTokenWithWrongKey(),
			expectedError: jwt.ErrSignatureInvalid,
		},
		{
			name:          "Test with malformed token",
			token:         "token",
			expectedError: jwt.ValidationError{Errors: jwt.ValidationErrorMalformed},
		},
		{
			name:          "Test with expired token",
			token:         makeToken(time.Now().Add(time.Minute * -5)),
			expectedError: jwt.ValidationError{Errors: jwt.ValidationErrorExpired},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateToken(mockContext, test.token)
			assert.Equal(t, test.expectedError, err)

			if err != nil {
				t.Logf("Failure: %s, error: %s", test.name, err)
			} else {
				t.Logf("Success: %s", test.name)
			}
		})
	}
}

func makeToken(expires time.Time) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(expires),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("secret_key"))
	return tokenString
}

func makeTokenWithWrongKey() string {
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("wrong_key"))
	return tokenString
}

