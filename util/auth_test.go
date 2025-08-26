package util

import (
	bytes "bytes"
	fmt "fmt"
	http "net/http"
	httptest "net/http/httptest"
	os "os"
	debug "runtime/debug"
	testing "testing"

	gin "github.com/gin-gonic/gin"
)

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error
*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		authorizationHeader string
		expectedError       bool
		validateTokenMock   func(ctx *gin.Context, token string) error
	}{
		{

			name:                "Valid Token",
			authorizationHeader: "Bearer valid-token",
			expectedError:       false,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return nil
			},
		},
		{

			name:                "Missing Authorization Header",
			authorizationHeader: "",
			expectedError:       true,
			validateTokenMock:   nil,
		},
		{

			name:                "Invalid Format Header",
			authorizationHeader: "random-text",
			expectedError:       true,
			validateTokenMock:   nil,
		},
		{

			name:                "Empty Authorization Header",
			authorizationHeader: "Bearer ",
			expectedError:       true,
			validateTokenMock:   nil,
		},
		{

			name:                "Invalid Token",
			authorizationHeader: "Bearer invalid-token",
			expectedError:       true,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return fmt.Errorf("token validation failed")
			},
		},
		{

			name:                "Header with Extra Whitespace",
			authorizationHeader: "Bearer          valid-token",
			expectedError:       false,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return nil
			},
		},
		{

			name:                "Non-Bearer Token Header",
			authorizationHeader: "Token valid-token",
			expectedError:       true,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return fmt.Errorf("unsupported scheme")
			},
		},
		{

			name:                "Extended Claims Token",
			authorizationHeader: "Bearer valid-token-with-claims",
			expectedError:       false,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return nil
			},
		},
		{

			name:                "Mixed-Case Bearer",
			authorizationHeader: "BEARER valid-token",
			expectedError:       false,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return nil
			},
		},
		{

			name:                "Multiple Tokens in Header",
			authorizationHeader: "Bearer valid-token extra-token",
			expectedError:       false,
			validateTokenMock: func(ctx *gin.Context, token string) error {
				return nil
			},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			buf := new(bytes.Buffer)
			writer := httptest.NewRecorder()
			_, _ = fmt.Fscanf(buf, "%s", os.Stdout)

			ctx, _ := gin.CreateTestContext(writer)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("authorization", tt.authorizationHeader)
			ctx.Request = req

			if tt.validateTokenMock != nil {
				err := tt.validateTokenMock(ctx, "mock-token")
				ctx.Set("ValidationError", err)
			}

			err := GetTokenInHeaderAndVerify(ctx)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error but received nil in test: %s", tt.name)
				} else {
					t.Logf("Expected error encountered: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error received in test: %s. Error: %v", tt.name, err)
				} else {
					t.Logf("Test '%s' passed successfully.", tt.name)
				}
			}
		})
	}
}
