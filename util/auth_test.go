package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/golang-jwt/jwt/v4"
)

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010
*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name          string
		token         string
		authorization string
		errExpected   bool
		errMessage    string
	}{
		{
			name:          "Valid Authorization Header and Token",
			token:         "Bearer valid-token",
			authorization: "",
			errExpected:   false,
			errMessage:    "",
		},
		{
			name:          "Missing Authorization Header",
			token:         "",
			authorization: "",
			errExpected:   true,
			errMessage:    "Authorization header is missing",
		},
		{
			name:          "Invalid Authorization Token",
			token:         "Bearer invalid-token",
			authorization: "",
			errExpected:   true,
			errMessage:    "Token is invalid",
		},
		{
			name:          "Malformatted Authorization Header",
			token:         "malformatted token",
			authorization: "",
			errExpected:   true,
			errMessage:    "malformatted authorization header",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			request, _ := http.NewRequest(http.MethodGet, "/", nil)

			if test.token != "" {
				request.Header.Set("Authorization", test.token)
			}

			context.Request = request
			err := GetTokenInHeaderAndVerify(context)

			if test.errExpected {
				assert.Error(t, err)
				assert.Equal(t, test.errMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02
*/
func TestValidateToken(t *testing.T) {

	parseWithClaimsOrig := jwt.ParseWithClaims
	jwt.ParseWithClaims = func(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
		return &jwt.Token{Valid: claims.(*Claims).Username != "invalid"}, nil
	}
	defer func() { jwt.ParseWithClaims = parseWithClaimsOrig }()

	testCases := []struct {
		name        string
		token       string
		mockClaims  *Claims
		expectError bool
		statusCode  int
	}{
		{
			name:        "Token is valid and parsable",
			token:       "valid_token",
			mockClaims:  &Claims{Username: "valid", RegisteredClaims: jwt.RegisteredClaims{}},
			expectError: false,
			statusCode:  http.StatusOK,
		},
		{
			name:        "Token is invalid",
			token:       "invalid_token",
			mockClaims:  &Claims{Username: "invalid", RegisteredClaims: jwt.RegisteredClaims{}},
			expectError: true,
			statusCode:  http.StatusUnauthorized,
		},
	}

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
