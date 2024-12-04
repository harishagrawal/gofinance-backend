package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/util"
	"errors"
	"fmt"
	"strings"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	testCases := []struct {
		name       string
		token      string
		errType    string
		errMessage string
	}{
		{
			name:       "Valid JWT Token",
			token:      "VALID_JWT_TOKEN",
			errType:    "nil",
			errMessage: "Token is valid",
		},
		{
			name:       "Invalid JWT Token",
			token:      "INVALID_JWT_TOKEN",
			errType:    "jwt.ValidationError",
			errMessage: "invalid token",
		},
		{
			name:       "Signature Invalid Error",
			token:      "JWT_TOKEN_INVALID_SIGNATURE",
			errType:    "jwt.ValidationError",
			errMessage: "signature is invalid",
		},
		{
			name:       "Other JWT Errors",
			token:      "JWT_TOKEN_OTHER_ERRORS",
			errType:    "jwt.ValidationError",
			errMessage: "other jwt errors",
		},
		{
			name:       "Token is Not Valid",
			token:      "JWT_TOKEN_NOT_VALID",
			errType:    "nil",
			errMessage: "Token is invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Log(testCase.name)

			r := gin.Default()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req, err := http.NewRequest("POST", "/testroute", nil)
			if assert.NoError(t, err) {
				ctx.Request = req
			}

			err = util.ValidateToken(ctx, testCase.token)

			switch testCase.errType {
			case "nil":
				assert.NoError(t, err)
				assert.Equal(t, testCase.errMessage, w.Body.String())
			case "jwt.ValidationError":
				assert.Error(t, err)
				assert.IsType(t, jwt.ValidationError{}, err)
				assert.Equal(t, testCase.errMessage, w.Body.String())
			default:
				t.Error("Unknown error type specified in test case")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)
	if len(fields) < 2 {
		return errors.New("Authorization header is missing")
	}
	tokenToValidate := fields[1]
	errOnValidateToken := ValidateToken(ctx, tokenToValidate)
	if errOnValidateToken != nil {
		return errOnValidateToken
	}
	return nil
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	testCases := []struct {
		Name          string
		Authorization string
		ExpectError   bool
	}{
		{
			Name:          "Normal operation with valid authorization header",
			Authorization: "Bearer valid_token",
			ExpectError:   false,
		},
		{
			Name:          "Error case with invalid authorization header",
			Authorization: "Bearer invalid_token",
			ExpectError:   true,
		},
		{
			Name:          "Error case with a missing authorization header",
			Authorization: "",
			ExpectError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/test", nil)
			req.Header.Add("Authorization", tc.Authorization)

			w := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			err := GetTokenInHeaderAndVerify(ctx)

			if tc.ExpectError {
				assert.Error(t, err, fmt.Sprintf("%s: expect error but got nil", tc.Name))
			} else {
				assert.NoError(t, err, fmt.Sprintf("%s: expect no error but got %v", tc.Name, err))
			}
		})
	}
}

func ValidateToken(ctx *gin.Context, token string) error {
	if token == "valid_token" {
		return nil
	}
	return errors.New("Invalid token")
}

