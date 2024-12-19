package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type TokenMock struct {
	mock.Mock
}
type Context struct {
	writermem	responseWriter
	Request		*http.Request
	Writer		ResponseWriter
	Params		Params
	handlers	HandlersChain
	index		int8
	fullPath	string
	engine		*Engine
	params		*Params
	skippedNodes	*[]skippedNode
	mu		sync.RWMutex
	Keys		map // This mutex protects Keys map.
	// Keys is a key/value pair exclusively for the context of each request.
	[string]any
	Errors		errorMsgs
	Accepted	[ // Errors is a list of errors attached to all the handlers/middlewares who used this context.
	// Accepted defines a list of manually accepted formats for content negotiation.
	]string
	queryCache	url.Values
	formCache	url.Values
	sameSite	http.SameSite
}// queryCache caches the query result from c.Request.URL.Query().
// SameSite allows a server to define a cookie attribute making it impossible for
// the browser to send this cookie along with cross-site requests.

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name             string
		mockAuthHeader   string
		mockValidateErr  error
		expectedErrorMsg string
	}{
		{
			name:             "Test with Valid authorization token in Header",
			mockAuthHeader:   "Bearer mockToken",
			mockValidateErr:  nil,
			expectedErrorMsg: "",
		},
		{
			name:             "Test with Invalid authorization token in Header",
			mockAuthHeader:   "Bearer mockToken",
			mockValidateErr:  errors.New("Invalid token"),
			expectedErrorMsg: "Invalid token",
		},
		{
			name:             "Test with Missing authorization token in Header",
			mockAuthHeader:   "",
			mockValidateErr:  nil,
			expectedErrorMsg: "index out of range",
		},
		{
			name:             "Test with Malformed authorization header in Header",
			mockAuthHeader:   "malformed_header",
			mockValidateErr:  nil,
			expectedErrorMsg: "index out of range",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			resp := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(resp)

			c.Request, _ = http.NewRequest("GET", "/", nil)
			c.Request.Header.Set("authorization", tt.mockAuthHeader)

			tokenMock := new(TokenMock)
			tokenMock.On("ValidateToken", c, strings.Fields(tt.mockAuthHeader)[1]).Return(tt.mockValidateErr)

			oldTokenValidator := ValidateToken
			ValidateToken = tokenMock.ValidateToken

			err := GetTokenInHeaderAndVerify(c)

			if tt.expectedErrorMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			}

			ValidateToken = oldTokenValidator
		})
	}
}

func (m *TokenMock) ValidateToken(ctx *gin.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func BaseContext() *gin.Context {
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)

	return ctx
}

func TestValidateToken(t *testing.T) {

	data := []struct {
		scenario     string
		token        string
		responseCode int
	}{
		{
			scenario:     "Valid Token",
			token:        createToken(),
			responseCode: http.StatusOK,
		},
		{
			scenario:     "Invalid Token Signature",
			token:        createTokenWithDifferentKey(),
			responseCode: http.StatusUnauthorized,
		},
		{
			scenario:     "Nonsensical Token",
			token:        "thisIsNotAToken",
			responseCode: http.StatusBadRequest,
		},
		{
			scenario:     "Expired Token",
			token:        createExpiredToken(),
			responseCode: http.StatusUnauthorized,
		},
		{
			scenario:     "Empty Token",
			token:        "",
			responseCode: http.StatusBadRequest,
		},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("%s", d.scenario), func(t *testing.T) {
			ctx := BaseContext()
			err := ValidateToken(ctx, d.token)
			if ctx.Writer.Status() != d.responseCode {
				t.Errorf("For scenario: %s, expected ReponseCode: %d, received ResponseCode: %d", d.scenario, d.responseCode, ctx.Writer.Status())
			}

			if err != nil && ctx.Writer.Status() == http.StatusOK {
				t.Errorf("Validation error but StatusOK.")
			} else if err == nil && ctx.Writer.Status() != http.StatusOK {
				t.Errorf("OK Validation but not StatusOK.")
			}
		})
	}
}

func createExpiredToken() string {
	return ""
}

func createToken() string {
	return ""
}

func createTokenWithDifferentKey() string {
	return ""
}

