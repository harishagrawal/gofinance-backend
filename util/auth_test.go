package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	util "github.com/wil-ckaew/gofinance-backend/util"
	"net/http/httptest"
	"testing"
	"net/http"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	var tests = []struct {
		scenario string
		token    string
		err      error
		status   int
	}{

		{"Valid Token", "valid_token", nil, 200},
		{"Invalid Token Signature", "invalid_signature_token", jwt.ErrSignatureInvalid, 401},
		{"Unreadable token - Malformed", "unreadable_token", jwt.ErrInvalidTokenFormat, 400},
		{"Invalid Token", "invalid_token", nil, 401},
	}

	for _, test := range tests {
		t.Logf("Running: %v", test.scenario)

		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		err := util.ValidateToken(c, test.token)

		if status := w.Code; status != test.status {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, test.status)
		}

		if err == nil && test.err != nil {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				err, test.err)
		}

		if err != nil && test.err != nil && err.Error() != test.err.Error() {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				err.Error(), test.err.Error())
		}

		fmt.Println("Response Code: ", w.Code)
		fmt.Println("Response Body: ", w.Body.String())
		fmt.Println()
	}
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	cases := []struct {
		name                string
		authorizationHeader string
		mockValidateToken   func(*gin.Context, string) error
		expectedError       bool
	}{
		{
			name:                "Successful Token Verification",
			authorizationHeader: "Bearer token",
			mockValidateToken: func(ctx *gin.Context, token string) error {

				return nil
			},
			expectedError: false,
		},
		{
			name:                "Missing Authorization Header",
			authorizationHeader: "",
			mockValidateToken: func(ctx *gin.Context, token string) error {
				return nil
			},
			expectedError: true,
		},
		{
			name:                "Invalid Token in Header",
			authorizationHeader: "Bearer invalidToken",
			mockValidateToken: func(ctx *gin.Context, token string) error {
				return jwt.NewValidationError("token is invalid", jwt.ValidationErrorMalformed)
			},
			expectedError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			resp := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(resp)
			req, _ := http.NewRequest("GET", "/somepath", nil)
			if tc.authorizationHeader != "" {
				req.Header["authorization"] = []string{tc.authorizationHeader}
			}
			c.Request = req

			validateToken = tc.mockValidateToken

			err := GetTokenInHeaderAndVerify(c)

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error %v but got %v", tc.expectedError, err != nil)
			}
		})
	}
}

