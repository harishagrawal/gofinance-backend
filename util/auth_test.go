package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

const InvalidToken = "Bearer invalid_token"
const ValidToken = "Bearer valid_token"func TestGetTokenInHeaderAndVerify(t *testing.T)
var getParser = func()/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func func TestGetTokenInHeaderAndVerify(t *testing.T) {

	testCases := []struct {
		name                string
		token               string
		authorizationHeader string
		expectedError       bool
	}{
		{
			name:                "Valid JWT token in Authorization header",
			token:               ValidToken,
			authorizationHeader: "authorization",
			expectedError:       false,
		},
		{
			name:                "Missing JWT token in Authorization header",
			token:               "",
			authorizationHeader: "authorization",
			expectedError:       true,
		},
		{
			name:                "Invalid JWT token in Authorization header",
			token:               InvalidToken,
			authorizationHeader: "authorization",
			expectedError:       true,
		},
		{
			name:                "JWT token not at the expected position in Authorization header",
			token:               "valid_token Bearer",
			authorizationHeader: "authorization",
			expectedError:       true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request, _ = http.NewRequest("POST", "/test", nil)
			c.Request.Header.Add(testCase.authorizationHeader, testCase.token)

			err := GetTokenInHeaderAndVerify(c)

			if testCase.expectedError {
				assert.NotNil(t, err, "Error should not be nil for case: %v", testCase.name)
			} else {
				assert.Nil(t, err, "Error should be nil for case: %v", testCase.name)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func func func(*jwt.Token) (interface{}, error)) (*jwt.Token, error) {
	args := m.Called(token, claims, keyfunc)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

func func TestValidateToken(t *testing.T) {
	testCases := []struct {
		name          string
		token         *MockJWT
		expectedError error
	}{
		{
			name: "Successful Token Validation",
			token: func() *MockJWT {
				mockToken := new(MockJWT)
				mockToken.On("ParseWithClaims", mock.Anything, mock.Anything, mock.Anything).Return(&jwt.Token{Valid: true}, nil)
				return mockToken
			}(),
			expectedError: nil,
		},
		{
			name: "Token with Replaced Signature",
			token: func() *MockJWT {
				mockToken := new(MockJWT)
				mockToken.On("ParseWithClaims", mock.Anything, mock.Anything, mock.Anything).Return(nil, jwt.ErrSignatureInvalid)
				return mockToken
			}(),
			expectedError: jwt.ErrSignatureInvalid,
		},
		{
			name: "Token with Invalid Structure",
			token: func() *MockJWT {
				mockToken := new(MockJWT)
				mockToken.On("ParseWithClaims", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("token contains an invalid number of segments"))
				return mockToken
			}(),
			expectedError: jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed),
		},
		{
			name: "Valid Token but Invalid",
			token: func() *MockJWT {
				mockToken := new(MockJWT)
				mockToken.On("ParseWithClaims", mock.Anything, mock.Anything, mock.Anything).Return(&jwt.Token{Valid: false}, nil)
				return mockToken
			}(),
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockContext := new(MockContext)
			switch tc.expectedError {
			case nil:
				mockContext.On("Next").Return()
			case jwt.ErrSignatureInvalid:
				mockContext.On("JSON", http.StatusUnauthorized, mock.Anything).Return()
			default:
				mockContext.On("JSON", http.StatusBadRequest, mock.Anything).Return()
			}

			err := ValidateToken(mockContext, tc.token)
			switch tc.expectedError {
			case nil:
				mockContext.AssertCalled(t, "Next")
				assert.NoError(t, err)
			case jwt.ErrSignatureInvalid:
				mockContext.AssertCalled(t, "JSON", http.StatusUnauthorized, mock.Anything)
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			default:
				mockContext.AssertCalled(t, "JSON", http.StatusBadRequest, mock.Anything)
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}

func func ValidateToken(ctx *gin.Context, tokenString string) error {
	parser := getParser()
	token, err := parser.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSignedKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return nil
	}

	ctx.Next()
	return nil
}

