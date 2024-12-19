package util

import (
	"errors"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
)

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func (m *MockedContext) GetHeader(s string) string {
	return m.mockedGetHeader
}

func (m *MockedValidator) ValidateToken(s string) error {
	args := m.Called(s)
	return args.Error(0)
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	testCases := []struct {
		name             string
		header           string
		mockValidateFunc func(m *MockedValidator)
		expectedError    error
	}{
		{
			name:   "Scenario 1: Valid Header and Token",
			header: "Token validToken",
			mockValidateFunc: func(m *MockedValidator) {
				m.On("ValidateToken", "validToken").Return(nil)
			},
			expectedError: nil,
		},
		{
			name:           "Scenario 2: Invalid Header",
			header:         "InvalidHeader",
			mockValidateFunc: func(m *MockedValidator) {},
			expectedError:  errors.New("no token present in the header"),
		},
		{
			name:   "Scenario 3: Invalid Token",
			header: "Token invalidToken",
			mockValidateFunc: func(m *MockedValidator) {
				m.On("ValidateToken", "invalidToken").
					Return(errors.New("invalid token"))
			},
			expectedError: errors.New("invalid token"),
		},
		{
			name:           "Scenario 4: Missing Authorization Header",
			header:         "",
			mockValidateFunc: func(m *MockedValidator) {},
			expectedError:  errors.New("no token present in the header"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		
			c := &gin.Context{}
			validator := new(MockedValidator)
			tc.mockValidateFunc(validator) 
			ctx := &MockedContext{
				Context:         *c,
				mockedGetHeader: tc.header,
			}

		
			ctx.Context.GetHeader = ctx.GetHeader
			
		
			err := GetTokenInHeaderAndVerify(&ctx.Context)

		
			if err != nil && tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, received error: %v", tc.expectedError, err)
				}
			} else if err != tc.expectedError {
				t.Errorf("Expected error: %v, received error: %v", tc.expectedError, err)
			}

			validator.AssertExpectations(t) 
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	secretKey := "secret_key"
	otherKey := "other_key"
	
	tests := []struct {
		name   string
		token  func() (string, error)
		error  error
		status int
	}{
		{
			name:   "Valid Token Provided",
			token:  func() (string, error) { return jwt.New(jwt.SigningMethodHS256).SignedString([]byte(secretKey)) },
			error:  nil,
			status: http.StatusOK,
		},
		{
			name:   "Invalid Token Provided",
			token:  func() (string, error) { return jwt.New(jwt.SigningMethodHS256).SignedString([]byte(otherKey)) },
			error:  jwt.ErrSignatureInvalid,
			status: http.StatusUnauthorized,
		},
		{
			name:   "Null Token Provided",
			token:  func() (string, error) { return "", nil },
			error:  jwt.ErrSignatureInvalid,
			status: http.StatusUnauthorized,
		},
		{
			name:   "Invalid Token Format Provided",
			token:  func() (string, error) { return "random_non_jwt_string", nil },
			error:  jwt.ErrSignatureInvalid,
			status: http.StatusUnauthorized,
		},
	
		{
			name:   "Expired Token Provided",
			token:  func() (string, error) { return "generate_expired_token_here", nil },
			error:  jwt.ErrExpired,
			status: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := tt.token()
			
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			err := ValidateToken(c, token)

			if err != tt.error {
				t.Errorf("expected error %v, got %v", tt.error, err)
			}
			
			if w.Code != tt.status {
				t.Errorf("expected status %d, got %d", tt.status, w.Code)
			}
		})
	}
}

