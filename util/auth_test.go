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

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	scenarios := []struct {
		name          string
		authorization string
		expectError   bool
	}{
		{
			name:          "Scenario 1: Enhancement of Token Verification",
			authorization: "Bearer abc",
			expectError:   false,
		},
		{
			name:          "Scenario 2: Empty Authorization Token Scenario",
			authorization: "",
			expectError:   true,
		},
		{
			name:          "Scenario 3: Test with Invalid Token",
			authorization: "Bearer invalid",
			expectError:   true,
		},
		{
			name:          "Scenario 4: No Token in 'authorization' header",
			authorization: "Bearer",
			expectError:   true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", s.authorization)
			c.Request = req

			err := GetTokenInHeaderAndVerify(c)
			if s.expectError {
				assert.NotNil(t, err, s.name)
			} else {
				assert.Nil(t, err, s.name)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

FUNCTION_DEF=func ValidateToken(ctx *gin.Context, token string) error 

 */
func TestValidateToken(t *testing.T) {

	testcases := []struct {
		name   string
		token  string
		status int
	}{
		{"Test Scenario 1: Valid JWT Token", "valid_token", http.StatusOK},
		{"Test Scenario 2: Invalid Signature in JWT Token", "faulty_token", http.StatusUnauthorized},
		{"Test Scenario 3: Token Parse Error", "malformed_token", http.StatusBadRequest},
		{"Test Scenario 4: Invalid JWT Token", "invalid_token", http.StatusUnauthorized},
		{"Test Scenario 5: No Token Provided", "", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := ValidateToken(ctx, tc.token)

			if err != nil {

				if jwtError, ok := err.(*jwt.ValidationError); ok {

					if uint32(tc.status) != jwtError.Errors {
						t.Errorf("expected '%v' but got '%v'", tc.status, jwtError.Errors)
					}
				} else {
					t.Error("Error is not jwt.ValidationError")
				}

			} else {
				if !ctx.IsAborted() {
					t.Errorf("expected status '%v' but got '%v'", tc.status, 200)
				}
			}
		})
	}
}

