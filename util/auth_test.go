package util_test

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"<YourModuleName>/util"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	gin.DefaultWriter = ioutil.Discard

	w := httptest.NewRecorder()
	r := gin.Default()
	c, _ := gin.CreateTestContext(w)

	validToken := "VALID_TOKEN"
	invalidToken := "INVALID_TOKEN"
	malformedToken := "MALFORMED_TOKEN"
	expiredToken := "EXPIRED_TOKEN"
	tokenTestCases := []struct {
		name        string
		token       string
		expectedErr error
	}{
		{"valid token", validToken, nil},
		{"invalid token", invalidToken, jwt.ErrSignatureInvalid},
		{"malformed token", malformedToken, errors.New("error decoding token")},
		{"expired token", expiredToken, errors.New("token is invalid")},
	}

	for _, testCase := range tokenTestCases {

		r.RouterGroup.Handle(fmt.Sprintf("/%s", testCase.name), func(c *gin.Context) {
			t.Run(testCase.name, func(t *testing.T) {
				err := ValidateToken(c, testCase.token)

				if err != testCase.expectedErr {
					t.Errorf("expected error %v, got %v", testCase.expectedErr, err)
				}
			})
		})
	}
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := gin.Default()

	testRouter.GET("/", func(c *gin.Context) {
		err := util.GetTokenInHeaderAndVerify(c)
		if err != nil {
			t.Fatalf("Expected no error but got: %v", err)
		}
	})

	w := performRequest(testRouter, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Add("Authorization", "Bearer token123")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

