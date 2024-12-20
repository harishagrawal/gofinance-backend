package util

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name              string
		requestHeader     string
		shouldReturnError bool
	}{
		{
			name:              "Valid Authorization Header",
			requestHeader:     "Bearer valid_token",
			shouldReturnError: false,
		},
		{
			name:              "Invalid Authorization Header Format",
			requestHeader:     "Invalid_Format",
			shouldReturnError: true,
		},
		{
			name:              "Invalid Token in Header",
			requestHeader:     "Bearer invalid_token",
			shouldReturnError: true,
		},
		{
			name:              "No Authorization Header in Request",
			requestHeader:     "",
			shouldReturnError: true,
		},
		{
			name:              "Empty Authorization Header Value",
			requestHeader:     "Bearer ",
			shouldReturnError: true,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Running test case: ", tt.name)
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)

			c.Request.Header.Set("Authorization", tt.requestHeader)
			err := GetTokenInHeaderAndVerify(c)

			if tt.shouldReturnError {
				assert.Error(t, err, "Error is expected")
			} else {
				assert.NoError(t, err, "Error is not expected")
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	var jwtSignedKey = []byte("your_jwt_secret")

	cases := []struct {
		name       string
		token      string
		wantStatus int
		wantError  string
	}{
		{
			name:       "Correct Token",
			token:      "",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid Token",
			token:      "invalid_token",
			wantStatus: http.StatusUnauthorized,
			wantError:  jwt.ErrSignatureInvalid.Error(),
		},
		{
			name:       "Null Token",
			token:      "",
			wantStatus: http.StatusUnauthorized,
			wantError:  jwt.ErrSignatureInvalid.Error(),
		},
		{
			name:       "Non-JWT Token",
			token:      "non_jwt_token",
			wantStatus: http.StatusBadRequest,
			wantError:  "token contains an invalid number of segments",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString(jwtSignedKey)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}
	cases[0].token = token

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/test", nil)
			response := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(response)
			context.Request = request
			err := ValidateToken(context, tt.token)

			if tt.wantError != "" {
				if assert.Error(t, err) {
					assert.Equal(t, tt.wantError, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantStatus, response.Code)
			if response.Code != http.StatusOK {
				assert.Equal(t, tt.wantError, response.Body.String())
			}
		})
	}
}

