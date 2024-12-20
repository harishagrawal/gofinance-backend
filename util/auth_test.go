package util

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
)


var jwtSignedKey = []byte("secret_key")/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func func TestGetTokenInHeaderAndVerify(t *testing.T) {
	tests := []struct {
		name                       string
		mockGetHeader              string
		mockAuthorizationHeaderKey string
		expectedError              string
		successWord                string
		capturedErrorWord          string
	}{
		{
			name:                       "Valid Token Test",
			mockGetHeader:              "Bearer testtoken",
			mockAuthorizationHeaderKey: "authorization",
			expectedError:              "",
			successWord:                "Valid token successfully verified",
			capturedErrorWord:          "Error: Valid token test failed",
		},
		{
			name:                       "Malformed Token Test",
			mockGetHeader:              "Bearer invalidtoken",
			mockAuthorizationHeaderKey: "authorization",
			expectedError:              "invalid token",
			successWord:                "Successfully handled malformed token",
			capturedErrorWord:          "Error: Malformed token test failed",
		},
		{
			name:                       "Missing Token Test",
			mockGetHeader:              "Bearer",
			mockAuthorizationHeaderKey: "authorization",
			expectedError:              "token missing",
			successWord:                "Successfully handled missing token case",
			capturedErrorWord:          "Error: Missing token test failed",
		},
		{
			name:                       "Invalid Token Test",
			mockGetHeader:              "Bearer invalidtoken",
			mockAuthorizationHeaderKey: "authorization",
			expectedError:              "token not valid",
			successWord:                "Successfully handled invalid token case",
			capturedErrorWord:          "Error: Invalid token test failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			mockContext := new(MockContext)
			mockContext.Context = c
			mockContext.On("GetHeader", test.mockAuthorizationHeaderKey).Return(test.mockGetHeader)

			err := GetTokenInHeaderAndVerify(mockContext)

			if len(test.expectedError) > 0 {

				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), test.expectedError)
				t.Log(test.capturedErrorWord)
			} else {

				assert.Nil(t, err)
				t.Log(test.successWord)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func func TestValidateToken(t *testing.T) {
	correctToken, _ := jwt.New(jwt.GetSigningMethod("HS256")).SignedString(jwtSignedKey)

	scenarios := []tokenScenarios{
		{
			description:   "Successful Token Validation",
			token:         correctToken,
			createContext: true,
			expectedError: nil,
			expectedCode:  http.StatusOK,
		},
		{
			description:   "Invalid Token Signature",
			token:         "invalid_token",
			createContext: true,
			expectedError: jwt.ErrSignatureInvalid,
			expectedCode:  http.StatusUnauthorized,
		},
		{
			description:   "Invalid Token Format",
			token:         "invalid_format",
			createContext: true,
			expectedError: jwt.NewValidationError("malformed", jwt.ValidationErrorMalformed),
			expectedCode:  http.StatusBadRequest,
		},
		{
			description:   "Missing Token",
			token:         "",
			createContext: true,
			expectedError: errors.New("No token present in request"),
			expectedCode:  http.StatusBadRequest,
		},
	}

	for _, s := range scenarios {
		t.Run(s.description, func(t *testing.T) {
			var ctx *gin.Context
			if s.createContext {
				r := gin.Default()
				ctx = &gin.Context{
					Request: httptest.NewRequest("GET", "/test", nil),
					Writer:  gin.DefaultWriter,
					Engine:  r,
				}
				ctx.Request.Header.Add("Authorization", "Bearer "+s.token)
			}

			err := ValidateToken(ctx, s.token)
			if err != nil {
				if err.Error() != s.expectedError.Error() {
					t.Errorf("failed %s: received %+v but expected %+v", s.description, err.Error(), s.expectedError.Error())
				}
				resp := ctx.Writer.(*gin.response).Result()
				if resp.StatusCode != s.expectedCode {
					t.Errorf("failed %s: received status code %d but expected %d", s.description, resp.StatusCode, s.expectedCode)
				}
			} else {
				if s.expectedError != nil {
					t.Errorf("failed %s: received no error but expected %+v", s.description, s.expectedError)
				}
			}
			t.Logf("success %s", s.description)
		})
	}
}

