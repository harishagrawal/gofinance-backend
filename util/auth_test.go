package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
)


var getTokenInHeaderTestData = []struct {
/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	for _, test := range getTokenInHeaderTestData {
		t.Run(test.name, func(t *testing.T) {
			ctx := generateMockContext(test.authorization)
			err := GetTokenInHeaderAndVerify(ctx)

			if (err == nil) == test.expectedResult {
				t.Logf("PASS: %s", test.name)
			} else {
				t.Errorf("FAIL: %s", test.name)
			}

			assert.IsType(t, http.ErrNotSupported, err)

			if err != nil {
				assert.Equal(t, "Token is invalid", err.Error())
			}
		})
	}
}

func generateMockContext(authorization string) *gin.Context {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("authorization", authorization)

	return &gin.Context{
		Request: req,
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func GenerateTokenWithIncorrectSignature() string {

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte(base64.StdEncoding.EncodeToString([]byte("Invalid_secret_key"))))
	if err != nil {
		return ""
	}

	return tokenString
}

func GenerateValidToken() string {

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString(jwtSignedKey)
	if err != nil {
		return ""
	}

	return tokenString
}

func TestValidateToken(t *testing.T) {

	testCases := []struct {
		name          string
		token         string
		expectedError error
	}{
		{"Test ValidateToken With Valid Token", GenerateValidToken(), nil},
		{"Test ValidateToken With Invalid Token", "invalidToken", jwt.ErrSignatureInvalid},
		{"Test ValidateToken With Empty Token", "", jwt.ErrSignatureInvalid},
		{"Test ValidateToken With Incorrect Signature", GenerateTokenWithIncorrectSignature(), jwt.ErrSignatureInvalid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			err := ValidateToken(c, tc.token)

			if err != tc.expectedError {
				t.Errorf("Test case '%s' failed, expected error '%s', but got '%s'", tc.name, tc.expectedError, err)
			} else {
				t.Logf("Test case '%s' passed successfully", tc.name)
			}
		})
	}
}

