package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"errors"
	"fmt"
	"strings"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		expResult string
		expStatus int
	}{
		{
			name:      "scenario 1: valid token",
			token:     createTestToken(),
			expResult: "",
			expStatus: http.StatusOK,
		},
		{
			name:      "scenario 2: invalid signature",
			token:     createTestToken() + "A",
			expResult: "signature is invalid",
			expStatus: http.StatusUnauthorized,
		},
		{
			name:      "scenario 3: token is empty",
			token:     "",
			expResult: "signature is invalid",
			expStatus: http.StatusBadRequest,
		},
		{
			name:      "scenario 4: invalid token",
			token:     "randomText",
			expResult: "token contains an invalid number of segments",
			expStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)

			err := ValidateToken(ctx, tc.token)

			if err != nil && err.Error() != tc.expResult {
				t.Fatalf("Expected error %v, got error %v instead", tc.expResult, err)
			}

			if tc.expStatus != recorder.Code {
				t.Fatalf("Expected HTTP status %v, got %v", tc.expStatus, recorder.Code)
			}
		})
	}
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}

	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
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

	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return nil
	}

	ctx.Next()
	return nil
}

func createTestToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: "testUser",
	})

	ss, _ := token.SignedString(jwtSignedKey)
	return ss
}

/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func GetTokenInHeaderAndVerifyUsingMock(ctx *gin.Context, validateTokenFunc func(ctx *gin.Context, token string) error) error {
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)

	if len(fields) < 2 {
		return errors.New("Token is missing")
	}

	tokenToValidate := fields[1]
	errOnValidateToken := validateTokenFunc(ctx, tokenToValidate)

	if errOnValidateToken != nil {
		return errOnValidateToken
	}
	return nil
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {

	var mockValidateTokenErr error

	validateTokenMockFunc := func(ctx *gin.Context, token string) error {

		return mockValidateTokenErr
	}

	testCases := []struct {
		name                     string
		authorizationHeaderValue string
		wantError                bool
		expectValidateTokenError error
	}{
		{
			name:                     "Scenario 1: Successful Token Validation",
			authorizationHeaderValue: "Bearer validToken",
			wantError:                false,
			expectValidateTokenError: nil,
		},
		{
			name:                     "Scenario 2: Missing Authorization Header",
			authorizationHeaderValue: "",
			wantError:                true,
			expectValidateTokenError: nil,
		},
		{
			name:                     "Scenario 3: Invalid Token Validation",
			authorizationHeaderValue: "Bearer invalidToken",
			wantError:                true,
			expectValidateTokenError: errors.New("invalid token"),
		},
		{
			name:                     "Scenario 4: Authorization Header without Token",
			authorizationHeaderValue: "Bearer",
			wantError:                true,
			expectValidateTokenError: nil,
		},
	}

	for i, testCase := range testCases {
		t.Logf(" Running test case %d: %s\n", i+1, testCase.name)

		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("authorization", testCase.authorizationHeaderValue)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = req
		mockValidateTokenErr = testCase.expectValidateTokenError

		err := GetTokenInHeaderAndVerifyUsingMock(c, validateTokenMockFunc)

		if testCase.wantError {
			assert.Error(t, err, fmt.Sprintf("Expected error for test case %d: %s\n", i+1, testCase.name))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Did not expect error for test case %d: %s\n", i+1, testCase.name))
		}
	}
}

