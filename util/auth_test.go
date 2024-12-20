package util

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"time"
)


var jwtSignedKey = []byte("secret_key")
/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	var jwtSignedKey = []byte("secret_key")

	testCases := []struct {
		description     string
		givenAuthHeader string
		expectError     bool
	}{
		{
			description:     "Success - valid token",
			givenAuthHeader: "Bearer " + generateTokenWithSecret(jwtSignedKey),
			expectError:     false,
		},
		{
			description:     "Failure - Invalid token",
			givenAuthHeader: "Bearer invalid-token",
			expectError:     true,
		},
		{
			description:     "Failure - Empty token",
			givenAuthHeader: "",
			expectError:     true,
		},
		{
			description:     "Failure - Missing Authorization Header",
			givenAuthHeader: "",
			expectError:     true,
		},
		{
			description:     "Failure - Token with incorrect format",
			givenAuthHeader: "Bearer " + generateTokenWithIncorrectFormat(),
			expectError:     true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {

			w := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
			ctx.Request.Header.Set("authorization", tt.givenAuthHeader)

			err := GetTokenInHeaderAndVerify(ctx)

			if tt.expectError {
				assert.NotNil(t, err)
				t.Log("Error is not nil as expected: ", err.Error())
			} else {
				assert.Nil(t, err)
				t.Log("Error is nil as expected")
			}
		})
	}
}

func generateTokenWithIncorrectFormat() string {
	return "invalid.token.format"
}

func generateTokenWithSecret(secret []byte) string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenStr, _ := token.SignedString(secret)

	return tokenStr
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	tests := []struct {
		name         string
		token        string
		wantHTTPCode int
		wantErrMsg   string
	}{
		{
			name: "Valid Token",
			token: prepareToken(jwt.SigningMethodHS256, &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour).Unix(),
				},
			}, jwtSignedKey),
			wantHTTPCode: http.StatusOK,
		},
		{
			name: "Invalid Signature",
			token: prepareToken(jwt.SigningMethodHS256, &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour).Unix(),
				},
			}, []byte("invalid_secret_key")),
			wantHTTPCode: http.StatusUnauthorized,
		},
		{
			name:         "Invalid Token",
			token:        "invalid_token",
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name: "Expired Token",
			token: prepareToken(jwt.SigningMethodHS256, &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(-time.Hour).Unix(),
				},
			}, jwtSignedKey),
			wantHTTPCode: http.StatusUnauthorized,
		},
		{
			name:         "Null Token",
			token:        "",
			wantHTTPCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(&MockResponseWriter{})
			c.Request = &http.Request{}

			err := ValidateToken(c, tt.token)

			if assert.Error(t, err) {
				assert.Equal(t, tt.wantHTTPCode, c.Writer.Status())
				if tt.wantErrMsg != "" {
					assert.Contains(t, err.Error(), tt.wantErrMsg)
				}
			} else {
				assert.Equal(t, tt.wantHTTPCode, c.Writer.Status())
			}
		})
	}
}

func prepareToken(method jwt.SigningMethod, claims jwt.Claims, key []byte) string {
	token := jwt.NewWithClaims(method, claims)
	str, _ := token.SignedString(key)
	return str
}

