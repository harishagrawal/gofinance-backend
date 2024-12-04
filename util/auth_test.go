package util

import (
	"testing"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	t.Run("Scenario 1: Validate a Valid Token", func(t *testing.T) {
		token, ctx := setup()

		err := ValidateToken(ctx, token)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	t.Run("Scenario 2: Validate an Invalid Token", func(t *testing.T) {

		_, ctx := setup()

		err := ValidateToken(ctx, "invalid_token")

		if err != jwt.ErrSignatureInvalid {
			t.Errorf("Expected unauthorized error, got %v", err)
		}
	})

	t.Run("Scenario 3: Validate a Token with Malformed Payload", func(t *testing.T) {
		_, ctx := setup()

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.EQJrijcuhvE-fedORxZxliNCLS_gTnukkGvP_OhYi10"
		err := ValidateToken(ctx, token)

		if err == nil || !strings.Contains(err.Error(), "token contains an invalid number of segments") {
			t.Errorf("Expected bad request error, got %v", err)
		}

	})

	t.Run("Scenario 4: Validate an Empty Token", func(t *testing.T) {
		_, ctx := setup()

		err := ValidateToken(ctx, "")

		if err == nil || err != jwt.ErrSignatureInvalid {
			t.Errorf("Expected bad request error, got %v", err)
		}
	})
}

func setup() (string, *gin.Context) {
	secretKey := "secret_key"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: "testUser",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	})

	tokenString, _ := token.SignedString([]byte(secretKey))

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	return tokenString, c
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	tests := []struct {
		name     string
		prepare  func() (*gin.Context, string)
		validate func(*testing.T, error, string)
	}{
		{
			name: "Valid Token",
			prepare: func() (*gin.Context, string) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				token := "Bearer VALID_TOKEN"
				c.Request.Header.Set("authorization", token)
				return c, token
			},
			validate: func(t *testing.T, err error, _ string) {
				if err != nil {
					t.Errorf("Error while testing with valid token: %v", err)
				}
			},
		},
		{
			name: "Invalid Token",
			prepare: func() (*gin.Context, string) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				token := "Bearer INVALID_TOKEN"
				c.Request.Header.Set("authorization", token)
				return c, token
			},
			validate: func(t *testing.T, err error, _ string) {
				if err == nil {
					t.Errorf("Error was expected but we got no error")
				}
			},
		},
		{
			name: "Missing Token",
			prepare: func() (*gin.Context, string) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request.Header.Set("authorization", "")
				return c, ""
			},
			validate: func(t *testing.T, err error, _ string) {
				if err == nil {
					t.Errorf("Error was expected but we got no error")
				}
			},
		},
		{
			name: "Null Context",
			prepare: func() (*gin.Context, string) {
				return nil, ""
			},
			validate: func(t *testing.T, err error, _ string) {
				if err == nil {
					t.Errorf("Error was expected but we got no error")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, token := tt.prepare()
			err := GetTokenInHeaderAndVerify(ctx)
			tt.validate(t, err, token)
		})
	}
}

