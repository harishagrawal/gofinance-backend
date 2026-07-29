package util

import (
	http "net/http"
	strings "strings"
	testing "testing"
	errors "errors"
	debug "runtime/debug"
	gin "github.com/gin-gonic/gin"
)



var validateTokenFuncForTest = ValidateToken

type testCase struct {
	name               string
	headerKey          string
	headerValue        string
	ctxInit            func() *gin.Context
	validateTokenFunc  func(ctx *gin.Context, token string) error
	expectErr          bool
	expectPanic        bool
	expectToken        string
	expectErrString    string
}


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error 

*/
func GetTokenInHeaderAndVerifyForTest(ctx *gin.Context) error {
	if ctx == nil {
		panic("ctx is nil passed to GetTokenInHeaderAndVerify")
	}
	authorizationHeaderKey := ctx.GetHeader("authorization")
	if strings.TrimSpace(authorizationHeaderKey) == "" {
		panic("authorization header missing")
	}
	fields := strings.Fields(authorizationHeaderKey)
	if len(fields) < 2 {
		panic("token not found in header")
	}
	tokenToValidate := fields[1]
	errOnValiteToken := getTokenValidateProxy(ctx, tokenToValidate)
	if errOnValiteToken != nil {
		return errOnValiteToken
	}
	return nil
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []testCase{
		{
			name:        "Scenario 1: Successful Token Extraction and Validation",
			headerKey:   "authorization",
			headerValue: "Bearer validtoken",
			expectErr:   false,
			expectPanic: false,
			expectToken: "validtoken",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "validtoken" {
					t.Errorf("Expected token 'validtoken', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer validtoken")
				return c
			},
		},
		{
			name:            "Scenario 2: Authorization Header Missing",
			headerKey:       "",
			headerValue:     "",
			expectErr:       true,
			expectPanic:     true,
			expectErrString: "",
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				return c
			},
			validateTokenFunc: ValidateToken,
		},
		{
			name:        "Scenario 3: Authorization Header With No Token",
			headerKey:   "authorization",
			headerValue: "Bearer",
			expectErr:   true,
			expectPanic: true,
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer")
				return c
			},
			validateTokenFunc: ValidateToken,
		},
		{
			name:        "Scenario 4: Malformed Authorization Header (Multiple Words, But Wrong Format)",
			headerKey:   "authorization",
			headerValue: "Basic notabearertoken",
			expectErr:   false,
			expectPanic: false,
			expectToken: "notabearertoken",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "notabearertoken" {
					t.Errorf("Expected token 'notabearertoken', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Basic notabearertoken")
				return c
			},
		},
		{
			name:            "Scenario 5: ValidateToken Returns an Error",
			headerKey:       "authorization",
			headerValue:     "Bearer errorToken",
			expectErr:       true,
			expectPanic:     false,
			expectToken:     "errorToken",
			expectErrString: "token invalid",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "errorToken" {
					t.Errorf("Expected token 'errorToken', got %v", token)
				}
				return errors.New("token invalid")
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer errorToken")
				return c
			},
		},
		{
			name:        "Scenario 6: Authorization Header With Extra Whitespace",
			headerKey:   "authorization",
			headerValue: "Bearer    tokenvalue",
			expectErr:   false,
			expectPanic: false,
			expectToken: "tokenvalue",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "tokenvalue" {
					t.Errorf("Expected token 'tokenvalue', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer    tokenvalue")
				return c
			},
		},
		{
			name:        "Scenario 7: Authorization Header With More Than Two Elements",
			headerKey:   "authorization",
			headerValue: "Bearer tokenvalue extrapart",
			expectErr:   false,
			expectPanic: false,
			expectToken: "tokenvalue",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "tokenvalue" {
					t.Errorf("Expected token 'tokenvalue', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer tokenvalue extrapart")
				return c
			},
		},
		{
			name:        "Scenario 8: Authorization Header With Empty String",
			headerKey:   "authorization",
			headerValue: "",
			expectErr:   true,
			expectPanic: true,
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "")
				return c
			},
			validateTokenFunc: ValidateToken,
		},
		{
			name:        "Scenario 9: Authorization Header has only Whitespaces",
			headerKey:   "authorization",
			headerValue: "    ",
			expectErr:   true,
			expectPanic: true,
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "    ")
				return c
			},
			validateTokenFunc: ValidateToken,
		},
		{
			name:        "Scenario 10: Authorization Header Case Sensitivity",
			headerKey:   "Authorization",
			headerValue: "Bearer validtoken",
			expectErr:   false,
			expectPanic: false,
			expectToken: "validtoken",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "validtoken" {
					t.Errorf("Expected token 'validtoken', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("Authorization", "Bearer validtoken")
				return c
			},
		},
		{
			name:        "Scenario 11: Extremely Long Authorization Header Value",
			headerKey:   "authorization",
			headerValue: "Bearer " + strings.Repeat("A", 4096),
			expectErr:   false,
			expectPanic: false,
			expectToken: strings.Repeat("A", 4096),
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if len(token) != 4096 || token != strings.Repeat("A", 4096) {
					t.Errorf("Expected token of length 4096, got %v", len(token))
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer "+strings.Repeat("A", 4096))
				return c
			},
		},
		{
			name:        "Scenario 12: Token Contains Internal White Spaces",
			headerKey:   "authorization",
			headerValue: "Bearer abc def",
			expectErr:   false,
			expectPanic: false,
			expectToken: "abc",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				if token != "abc" {
					t.Errorf("Expected token 'abc', got %v", token)
				}
				return nil
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer abc def")
				return c
			},
		},
		{
			name:        "Scenario 13: Gin Context is Nil",
			expectErr:   true,
			expectPanic: true,
			ctxInit: func() *gin.Context {
				return nil
			},
			validateTokenFunc: ValidateToken,
		},
		{
			name:        "Scenario 14: ValidateToken Panics",
			headerKey:   "authorization",
			headerValue: "Bearer willpanic",
			expectErr:   true,
			expectPanic: true,
			expectToken: "willpanic",
			validateTokenFunc: func(ctx *gin.Context, token string) error {
				panic("validate token panic simulation")
			},
			ctxInit: func() *gin.Context {
				c, _ := gin.CreateTestContext(nil)
				c.Request = &http.Request{Header: make(http.Header)}
				c.Request.Header.Set("authorization", "Bearer willpanic")
				return c
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Errorf("Unexpected panic for scenario '%s': %v", tc.name, r)
					} else {
						t.Logf("Scenario '%s': Panic expected and occurred", tc.name)
					}
					t.Fail()
				}
			}()

			t.Logf("Scenario: %s", tc.name)

			validateTokenFuncForTest = tc.validateTokenFunc

			ctx := tc.ctxInit()
			if ctx == nil && tc.name != "Scenario 13: Gin Context is Nil" {
				t.Errorf("ctxInit returned nil context for non-nil scenario")
			}

			var err error
			if !tc.expectPanic {
				err = GetTokenInHeaderAndVerifyForTest(ctx)
			} else {
				GetTokenInHeaderAndVerifyForTest(ctx)
			}

			if tc.expectErr && err == nil && !tc.expectPanic {
				t.Errorf("Scenario '%s': Expected error, got nil", tc.name)
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Scenario '%s': Expected nil, got error: %v", tc.name, err)
			}
			if tc.expectErrString != "" && err != nil && !strings.Contains(err.Error(), tc.expectErrString) {
				t.Errorf("Scenario '%s': Error returned does not match expectation. Want substring: '%s', got: '%v'", tc.name, tc.expectErrString, err)
			}

			t.Logf("Scenario '%s': Completed: expectErr=%v, expectPanic=%v", tc.name, tc.expectErr, tc.expectPanic)
		})
	}
}

func getTokenValidateProxy(ctx *gin.Context, token string) error {
	return validateTokenFuncForTest(ctx, token)
}

