package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
func TestGetLocalTokenInHeaderAndVerify(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() string
		expected error
	}{
		{

			name: "Valid token",
			setup: func() string {
				return "Bearer " + generateAndValidateToken(false)
			},
			expected: nil,
		},
		{

			name: "Missing authorization header",
			setup: func() string {
				return ""
			},
			expected: errors.New("Token is invalid"),
		},
		{

			name: "Invalid token",
			setup: func() string {
				return "Bearer " + generateAndValidateToken(true)
			},
			expected: errors.New("Token is invalid"),
		},
		{

			name: "No token in Authorization Header",
			setup: func() string {
				return "Bearer "
			},
			expected: errors.New("Token is invalid"),
		},
		{

			name: "No Bearer Scheme in Authorization Header",
			setup: func() string {
				return generateAndValidateToken(false)
			},
			expected: errors.New("Token is invalid"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
			c.Request.Header.Set("Authorization", test.setup())

			err := getLocalTokenInHeaderAndVerify(c)

			if test.expected != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expected.Error(), err.Error())
				return
			}
			assert.Nil(t, err)
		})
	}
}

func generateAndValidateToken(invalid bool) string {
	if invalid {

		return "invalid_token"
	}

	return "valid_token"
}

func getLocalTokenInHeaderAndVerify(c *gin.Context) error {
	authorizationHeaderKey := c.GetHeader("Authorization")
	if authorizationHeaderKey == "Bearer valid_token" {
		return nil
	}
	return errors.New("Token is invalid")
}

