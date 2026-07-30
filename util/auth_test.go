package util

import (
	http "net/http"
	testing "testing"
	bytes "bytes"
	fmt "fmt"
	debug "runtime/debug"
	gin "github.com/gin-gonic/gin"
)



var validateTokenTestCB func(ctx *gin.Context, token string) error




/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	captureOut := func(fn func()) string {
		var buf bytes.Buffer
		old := gin.DefaultWriter
		gin.DefaultWriter = &buf
		defer func() { gin.DefaultWriter = old }()
		fn()
		return buf.String()
	}

	type testCase struct {
		name            string
		authHeaderKey   string
		authHeaderValue string
		mockValidateErr error
		expectErr       bool
		expectPanic     bool
		validateTokenCB func(*gin.Context, string) error
	}

	validToken := "valid.jwt.token"
	invalidToken := "invalid.jwt.token"
	nonPrintableToken := "\x00\x01\x02binarytoken"

	tests := []testCase{
		{
			name:            "Scenario 1: Properly formatted Bearer, valid token",
			authHeaderKey:   "authorization",
			authHeaderValue: "Bearer " + validToken,
			mockValidateErr: nil,
			expectErr:       false,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token != validToken {
					return fmt.Errorf("unexpected token value")
				}
				return nil
			},
		},
		{
			name:            "Scenario 2: Bearer token present, but token fails validation",
			authHeaderKey:   "authorization",
			authHeaderValue: "Bearer " + invalidToken,
			mockValidateErr: fmt.Errorf("token invalid"),
			expectErr:       true,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token == invalidToken {
					return fmt.Errorf("token invalid")
				}
				return nil
			},
		},
		{
			name:            "Scenario 3: Missing authorization header",
			authHeaderKey:   "",
			authHeaderValue: "",
			expectErr:       true,
			expectPanic:     true,
		},
		{
			name:            "Scenario 4: Empty authorization header",
			authHeaderKey:   "authorization",
			authHeaderValue: "",
			expectErr:       true,
			expectPanic:     true,
		},
		{
			name:            "Scenario 5: Malformed (only 'Bearer' no token) header",
			authHeaderKey:   "authorization",
			authHeaderValue: "Bearer",
			expectErr:       true,
			expectPanic:     true,
		},
		{
			name:            "Scenario 6: Whitespaces in authorization header",
			authHeaderKey:   "authorization",
			authHeaderValue: "    ",
			expectErr:       true,
			expectPanic:     true,
		},
		{
			name:            "Scenario 7: Non-Bearer (Basic <token>) authorization type",
			authHeaderKey:   "authorization",
			authHeaderValue: "Basic " + validToken,
			mockValidateErr: nil,
			expectErr:       false,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token != validToken {
					return fmt.Errorf("wrong token extracted")
				}
				return nil
			},
		},
		{
			name:            "Scenario 8: Authorization header with extra fields",
			authHeaderKey:   "authorization",
			authHeaderValue: "Bearer " + validToken + " something-extra",
			mockValidateErr: nil,
			expectErr:       false,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token != validToken {
					return fmt.Errorf("wrong token extracted from extra fields")
				}
				return nil
			},
		},
		{
			name:            "Scenario 9: Non-printable/binary data token in header",
			authHeaderKey:   "authorization",
			authHeaderValue: "Bearer " + nonPrintableToken,
			mockValidateErr: fmt.Errorf("binary token error"),
			expectErr:       true,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token == nonPrintableToken {
					return fmt.Errorf("binary token error")
				}
				return nil
			},
		},
		{
			name:            "Scenario 10: Authorization header with case variation",
			authHeaderKey:   "Authorization",
			authHeaderValue: "Bearer " + validToken,
			mockValidateErr: nil,
			expectErr:       false,
			expectPanic:     false,
			validateTokenCB: func(ctx *gin.Context, token string) error {
				if token != validToken {
					return fmt.Errorf("token value mismatch")
				}
				return nil
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test: %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Errorf("Scenario did not expect panic: got panic: %v", r)
					} else {
						t.Log("Expected panic occurred. Defensive programming or improved input validation recommended.")
					}
				}
				resetTestValidateToken()
			}()

			if tc.validateTokenCB != nil {
				setTestValidateToken(tc.validateTokenCB)
			}

			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(nil)
			req, _ := http.NewRequest("GET", "/", nil)
			if tc.authHeaderKey != "" {
				req.Header.Set(tc.authHeaderKey, tc.authHeaderValue)
			}
			c.Request = req

			var captured string
			ran := false

			captured = captureOut(func() {
				ran = true
				var err error

				err = GetTokenInHeaderAndVerify(c)
				t.Logf("Scenario: %s", tc.name)
				t.Logf("Authorization Header Key: %q", tc.authHeaderKey)
				t.Logf("Authorization Header Value: %q", tc.authHeaderValue)

				if tc.expectPanic {
					t.Logf("Test expects panic, verify t.Fail() will be triggered in defer.")
				} else if tc.expectErr {
					if err == nil {
						t.Errorf("Expected error, got nil. Diagnostic output:\n%s", captured)
					} else {
						t.Logf("Got expected error: %v", err)
					}
				} else {
					if err != nil {
						t.Errorf("Expected nil error, got: %v. Diagnostic output:\n%s", err, captured)
					} else {
						t.Logf("Success: nil error returned as expected.")
					}
				}
			})
			if !ran {
				t.Errorf("Test inner function did not run")
			}
		})
	}
}

func resetTestValidateToken() {
	validateTokenTestCB = nil
}

func setTestValidateToken(cb func(ctx *gin.Context, token string) error) {
	validateTokenTestCB = cb
}

