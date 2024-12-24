package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

type Claims struct{
	jwt.StandardClaims
}

var ValidateToken = func(ctx *gin.Context, token string) error {
	return validateTokenImpl(ctx, token)
}

func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHeaderKey := ctx.GetHeader("Authorization")
	fields := strings.Fields(authorizationHeaderKey)
	if len(fields) < 2 {
		return errors.New("no token provided in header")
	}
	tokenToValidate := fields[1]
	errOnValiteToken := ValidateToken(ctx, tokenToValidate)
	if errOnValiteToken != nil {
		return errOnValiteToken
	}
	return nil
}

func validateTokenImpl(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSignedKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return err
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return errors.New("Token is invalid")
	}

	ctx.Next()
	return nil
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	testCases := []struct {
		name                    string
		authorizationHeader     string
		mockValidateToken       func(*gin.Context, string) error
		expectedError           error
	}{
		{
			name:                    "Scenario 1: Successful Token Validation",
			authorizationHeader:     "Bearer valid-token",
			mockValidateToken:       func(ctx *gin.Context, token string) error { return nil },
			expectedError:           nil,
		},
		{
			name:                    "Scenario 2: Invalid Token in Authorization Header",
			authorizationHeader:     "Bearer invalid-token",
			mockValidateToken:       func(ctx *gin.Context, token string) error { return errors.New("invalid token") },
			expectedError:           errors.New("invalid token"),
		},
		{
			name:                    "Scenario 3: Absence of Token in Authorization Header",
			authorizationHeader:     "Bearer",
			mockValidateToken: 	  func(ctx *gin.Context, token string) error { return errors.New("no token provided in header") },
			expectedError:           errors.New("no token provided in header"),
		},
		{
			name:                    "Scenario 4: Absence of Authorization Header",
			authorizationHeader:     "",
			mockValidateToken: 	  func(ctx *gin.Context, token string) error { return errors.New("no authorization header provided") },
			expectedError:           errors.New("no authorization header provided"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockContext.Request, _ = http.NewRequest("POST", "/api/test", nil)
			mockContext.Request.Header.Add("Authorization", tc.authorizationHeader)

			oldValidateToken := ValidateToken
			defer func() {
				ValidateToken = oldValidateToken
			}()
			ValidateToken = tc.mockValidateToken

			err := GetTokenInHeaderAndVerify(mockContext)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
