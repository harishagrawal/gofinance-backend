package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

type ValidateTokenMock struct {
	ErrorToReturn error
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
func (m *ValidateTokenMock) Execute(ctx *gin.Context, token string) error {
	return m.ErrorToReturn
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {

	tests := []struct {
		name        string
		setupFunc   func() *gin.Context
		expectedErr error
	}{
		{
			name: "valid authorization token in Header",
			setupFunc: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("GET", "/token", nil)
				c.Request.Header.Set("Authorization", "Bearer valid_token")
				return c
			},
			expectedErr: nil,
		},
		{
			name: "Invalid authorization token in Header",
			setupFunc: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("GET", "/token", nil)
				c.Request.Header.Set("Authorization", "Bearer invalid_token")
				return c
			},
			expectedErr: errors.New("Invalid token"),
		},
		{
			name: "Missing authorization token in Header",
			setupFunc: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("GET", "/token", nil)
				return c
			},
			expectedErr: errors.New("missing authorization token in header"),
		},
		{
			name: "malformed authorization header in Header",
			setupFunc: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("GET", "/token", nil)
				c.Request.Header.Set("Authorization", "Bearer")
				return c
			},
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateTokenMock := &ValidateTokenMock{
				ErrorToReturn: tt.expectedErr,
			}
			err := GetTokenInHeaderAndVerify(tt.setupFunc(), validateTokenMock.Execute)
			if err != nil {
				if !strings.Contains(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("got error '%v', expected '%v'", err, tt.expectedErr)
				}
			} else if tt.expectedErr != nil {
				t.Errorf("got no error, but expected Err '%v'", tt.expectedErr)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	jwtSignedKey := []byte("secret_key")

	validToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString(jwtSignedKey)
	wrongKey := []byte("wrongSecret")
	invalidSignedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString(wrongKey)
	expiredTokenClaims := jwt.MapClaims{}
	expiredTokenClaims["exp"] = 0
	expiredToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredTokenClaims).SignedString(jwtSignedKey)

	testCases := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "Valid Token",
			token:   validToken,
			wantErr: false,
		},
		{
			name:    "Invalid Token Signature",
			token:   invalidSignedToken,
			wantErr: true,
		},
		{
			name:    "Nonsensical Token",
			token:   "thisIsNotAToken",
			wantErr: true,
		},
		{
			name:    "Expired Token",
			token:   expiredToken,
			wantErr: true,
		},
		{
			name:    "Empty Token",
			token:   "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			err := ValidateToken(c, tc.token)

			if tc.wantErr {
				require.Error(t, err, tc.name)
			} else {
				require.NoError(t, err, tc.name)
			}
		})
	}
}

