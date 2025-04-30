package util

import (
	errors "errors"
	http "net/http"
	httptest "net/http/httptest"
	testing "testing"

	gin "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type ginContextMock struct {
	mock.Mock
	*gin.Context
}
type ginResponseWriterMock struct {
	mock.Mock
	gin.ResponseWriter
}
type MockClaims struct {
	Username          string
	ValidationMessage string
	Token             string
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010

FUNCTION_DEF=func GetTokenInHeaderAndVerify(ctx *gin.Context) error
*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	scuttinario := []struct {
		name            string
		tokenInHeader   string
		mockValidateTok func(*gin.Context, string) error
		wantErr         bool
	}{
		{
			name:          "Successful token retrieval and validation",
			tokenInHeader: "Bearer validToken",
			mockValidateTok: func(*gin.Context, string) error {
				return nil
			},
			wantErr: false,
		},
		{
			name:          "Token not present",
			tokenInHeader: "",
			wantErr:       true,
		},
		{
			name:          "Token validation failure",
			tokenInHeader: "Bearer invalidToken",
			mockValidateTok: func(*gin.Context, string) error {
				return jwt.NewValidationError("", jwt.ValidationErrorUnverifiable)
			},
			wantErr: true,
		},
	}

	for _, tt := range scuttinario {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic encountered so failing test due to : %v\n", r)
					t.Fail()
				}
			}()

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest(http.MethodPost, "/", nil)
			req.Header.Add("Authorization", tt.tokenInHeader)
			ctx.Request = req

			err := GetTokenInHeaderAndVerify(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenInHeaderAndVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02

FUNCTION_DEF=func ValidateToken(ctx *gin.Context, token string) error
*/
func (m *MockClaims) Valid() error {
	if m.Token == m.Username {
		return nil
	}
	return errors.New(m.ValidationMessage)
}

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		err        error
		httpStatus int
		claims     jwt.Claims
	}{
		{
			name:       "Valid Token",
			token:      "validToken",
			err:        nil,
			httpStatus: http.StatusOK,
			claims:     &MockClaims{Username: "validToken", Token: "validToken", ValidationMessage: ""},
		},
		{
			name:       "Invalid Token",
			token:      "invalidToken",
			err:        errors.New("Invalid token"),
			httpStatus: http.StatusUnauthorized,
			claims:     &MockClaims{Username: "validToken", Token: "invalidToken", ValidationMessage: "Invalid token"},
		},
		{
			name:       "Malformed Token",
			token:      "malformedToken",
			err:        errors.New("Malformed token"),
			httpStatus: http.StatusBadRequest,
			claims:     &MockClaims{Username: "validToken", Token: "malformedToken", ValidationMessage: "Malformed token"},
		},
		{
			name:       "Token without Bearer schema",
			token:      "",
			err:        errors.New("Token without bearer schema"),
			httpStatus: http.StatusBadRequest,
			claims:     &MockClaims{Username: "validToken", Token: "", ValidationMessage: "Token without bearer schema"},
		},
		{
			name:       "Token with incorrect signing method",
			token:      "incorrectSigningMethodToken",
			err:        errors.New("Token with incorrect signing method"),
			httpStatus: http.StatusUnauthorized,
			claims:     &MockClaims{Username: "validToken", Token: "incorrectSigningMethodToken", ValidationMessage: "Token with incorrect signing method"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, TestValidateToken failed. %v", r)
					t.Fail()
				}
			}()

			ctx := &gin.Context{}

			token := test.token
			ctx.Set("token", token)

			err := ValidateToken(ctx, token)

			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			}

			status, exists := ctx.Get("status")
			if exists {
				assert.Equal(t, test.httpStatus, status)
			} else {
				t.Log("Status not set in context")
				t.Fail()
			}
		})
	}
}
