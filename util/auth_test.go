package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	newContext := func(req *http.Request) (*gin.Context, *gin.Engine) {
		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		c.Request = req
		return c, router
	}

	testCases := []struct {
		name              string
		token             string
		shouldReturnError bool
	}{
		{
			name:              "Valid Token",
			token:             "Bearer validtoken",
			shouldReturnError: false,
		},
		{
			name:              "Invalid Token",
			token:             "Bearer invalidtoken",
			shouldReturnError: true,
		},
		{
			name:              "No Token in Header",
			token:             "",
			shouldReturnError: true,
		},
		{
			name:              "Invalid Authorization Header Format",
			token:             "invalid header",
			shouldReturnError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("authorization", tc.token)

			c, _ := newContext(req)

			err = GetTokenInHeaderAndVerify(c)

			if (err != nil) != tc.shouldReturnError {
				t.Fatalf("expected error: %v, got: %v", tc.shouldReturnError, err != nil)
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
		name    string
		token   string
		isError bool
		error   error
	}{
		{
			name: "Validate Working Token",
		},
		{
			name: "Validate Token with Invalid Signature Error",

			isError: true,
			error:   jwt.ErrSignatureInvalid,
		},
		{
			name: "Validate Token with Bad Request Error",

			isError: true,

			error: jwt.NewValidationError("", jwt.ValidationErrorMalformed),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := gin.Default()
			router.Use(func(context *gin.Context) {
				err := ValidateToken(context, test.token)

				if test.isError {
					assert.Error(t, err)
					assert.Equal(t, test.error, err)
				} else {
					assert.NoError(t, err)
				}
			})

			router.GET(test.name, func(context *gin.Context) {})

			router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/"+test.name, nil))
		})
	}
}

