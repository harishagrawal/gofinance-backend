package util

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"net/http/httptest"
	"strings"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	var tokenTests = []struct {
		name           string
		token          string
		expectedError  error
		expectedStatus int
	}{
		{"Successful Validation of Valid Token", "valid token here with valid signature", nil, 200},
		{"Invalid Signature", "token with invalid signature", jwt.ErrSignatureInvalid, 401},
		{"Invalid Token", "invalid token format", jwt.ErrSignatureInvalid, 400},
		{"Token is not Valid", "token invalid but well formatted", errors.New("Token is invalid"), 401},
		{"Empty Token", "", jwt.ErrSignatureInvalid, 400},
	}

	for _, tt := range tokenTests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			mockGinContext := &MockGinContext{
				Context:  c,
				Response: w,
			}

			err := ValidateToken(mockGinContext, tt.token)

			if err != tt.expectedError {
				t.Errorf("Test %s failed, expected %v but got %v", tt.name, tt.expectedError, err)
			}

			if mockGinContext.Response.Code != tt.expectedStatus {
				t.Errorf("Test %s failed, expected HTTP status %v but got %v", tt.name, tt.expectedStatus, mockGinContext.Response.Code)
			}

			if tt.expectedError == nil {
				t.Logf("Test %s passed", tt.name)
			}

		})
	}
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	tests := []struct {
		name                       string
		authorizationHeaderKey     string
		mockValidateTokenReturnErr error
		expectError                bool
	}{
		{
			name:                       "Normal operation with valid input",
			authorizationHeaderKey:     "Bearer token",
			mockValidateTokenReturnErr: nil,
			expectError:                false,
		},
		{
			name:                       "Invalid operation with non-exist authorization header",
			authorizationHeaderKey:     "",
			mockValidateTokenReturnErr: nil,
			expectError:                true,
		},
		{
			name:                       "Invalid operation with invalid token",
			authorizationHeaderKey:     "Bearer invalid_token",
			mockValidateTokenReturnErr: jwt.NewValidationError("Validation error", jwt.ValidationErrorMalformed),
			expectError:                true,
		},
		{
			name:                       "Handle Edge case with empty authorization header",
			authorizationHeaderKey:     "Bearer ",
			mockValidateTokenReturnErr: jwt.NewValidationError("Validation error", jwt.ValidationErrorNotValidYet),
			expectError:                true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
			c.Request.Header.Add("authorization", tt.authorizationHeaderKey)

			mockJWTService := new(MockJWTService)
			mockJWTService.On("ValidateToken", c, strings.TrimPrefix(tt.authorizationHeaderKey, "Bearer ")).Return(tt.mockValidateTokenReturnErr)

			err := GetTokenInHeaderAndVerify(c)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockJWTService.AssertExpectations(t)
		})
	}
}

func (m *MockJWTService) ValidateToken(ctx *gin.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

