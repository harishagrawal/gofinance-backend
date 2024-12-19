package util

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http/httptest"
	"errors"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/your_username/your_repo/util"
)

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func CreateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "test"

	tokenString, _ := token.SigningString()

	return tokenString
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {

	tests := []struct {
		name   string
		header string
		token  string
		err    bool
	}{
		{name: "Authorization header present and JWT Token Valid", header: "authorization", token: CreateToken(), err: false},
		{name: "Authorization header not present", header: "", token: "", err: true},
		{name: "Authorization header present but Token Invalid", header: "authorization", token: "InvalidToken", err: true},
		{name: "JWT Token not present in the Authorization header", header: "authorization", token: "", err: true},
	}

	for i, tt := range tests {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		if tt.header != "" {

			c.Request.Header.Add(tt.header, "Bearer "+tt.token)
		}

		err := GetTokenInHeaderAndVerify(c)

		if (err != nil) != tt.err {
			t.Errorf("Test case %d: %s: Expected error %t but got error %v", i, tt.name, tt.err, err)
		} else {
			t.Logf("Test case %d: %s: passed as expected", i, tt.name)
		}
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	tests := []struct {
		name               string
		token              string
		mockContext        func() (*gin.Context, *httptest.ResponseRecorder)
		setup              func()
		wantErr            bool
		errExpected        error
		httpStatusExpected int
	}{
		{
			name:  "Valid Token Test",
			token: "valid_token",
			mockContext: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c, w
			},
			setup: func() {

			},
			wantErr:            false,
			errExpected:        nil,
			httpStatusExpected: http.StatusOK,
		},
		{
			name:  "Invalid Signature Token Test",
			token: "invalid_signature",
			mockContext: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c, w
			},
			setup:              nil,
			wantErr:            true,
			errExpected:        jwt.ErrSignatureInvalid,
			httpStatusExpected: http.StatusUnauthorized,
		},
		{
			name:  "Parse error Token Test",
			token: "unparseable_token",
			mockContext: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c, w
			},
			setup:              nil,
			wantErr:            true,
			errExpected:        errors.New("Token could not be parsed"),
			httpStatusExpected: http.StatusBadRequest,
		},
		{
			name:  "Invalid Token Test",
			token: "invalid_token",
			mockContext: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c, w
			},
			setup:              nil,
			wantErr:            true,
			errExpected:        errors.New("Token is invalid"),
			httpStatusExpected: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			c, w := tt.mockContext()
			err := util.ValidateToken(c, tt.token)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.errExpected, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.httpStatusExpected, w.Code)
		})
	}
}

