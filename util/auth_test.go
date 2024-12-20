package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var jwtSignedKey = []byte("secret_key")
/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func TestGetTokenInHeaderAndVerify(t *testing.T) {
	tests := []struct {
		name          string
		authorization string
		errExpect     error
	}{
		{"Successful Token Verification", "Bearer token", nil},
		{"Header Key Missing", "", errors.New("The authorization header is not presented")},
		{"Invalid Token Verification", "Bearer invalid_token", errors.New("Token is invalid")},
		{"Not Enough Fields In Header ", "Bearer ", errors.New("Length of splitted string by space is less than 2")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request.Header.Set("authorization", tt.authorization)
			err := GetTokenInHeaderAndVerify(ctx)
			if tt.errExpect != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.errExpect.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	w := httptest.NewRecorder()
	_ctx, _ := gin.CreateTestContext(w)

	validToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{}).SignedString(jwtSignedKey)

	invalidSignatureToken := validToken + "invalid-signature"
	invalidFormatToken := "invalid-format-token"

	testCases := map[string]struct {
		context    *gin.Context
		token      string
		err        error
		statusCode int
	}{
		"Valid Token Test": {
			context:    _ctx,
			token:      validToken,
			err:        nil,
			statusCode: http.StatusOK,
		},
		"Invalid Signature Test": {
			context:    _ctx,
			token:      invalidSignatureToken,
			err:        jwt.ErrSignatureInvalid,
			statusCode: http.StatusUnauthorized,
		},
		"Token String Format Error Test": {
			context:    _ctx,
			token:      invalidFormatToken,
			err:        jwt.NewValidationError("", jwt.ValidationErrorMalformed),
			statusCode: http.StatusBadRequest,
		},
		"Null Token Test": {
			context:    _ctx,
			token:      "",
			err:        jwt.ErrSignatureInvalid,
			statusCode: http.StatusUnauthorized,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			err := ValidateToken(test.context, test.token)

			if err != nil {
				assert.Equal(t, test.err, err)
			}

			assert.Equal(t, test.statusCode, w.Code)
		})
	}
}

