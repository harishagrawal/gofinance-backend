package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var jwtSignedKey = []byte("jwt_token")
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name       string
		token      string
		statusCode int
		expectErr  bool
	}{
		{
			name:       "Valid Token",
			token:      "Bearer " + createTokenWithUsername("test"),
			statusCode: http.StatusOK,
			expectErr:  false,
		},
		{
			name:       "Invalid Token",
			token:      "Bearer invalid_token",
			statusCode: http.StatusUnauthorized,
			expectErr:  true,
		},
		{
			name:       "Empty Token",
			token:      "",
			statusCode: http.StatusUnauthorized,
			expectErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tc.token)
			resp := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(resp)
			c.Request = req

			err := GetTokenInHeaderAndVerify(c)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.statusCode, resp.Code)
		})
	}
}

func createTokenWithUsername(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
	})

	tokenString, _ := token.SignedString(jwtSignedKey)
	return tokenString
}

