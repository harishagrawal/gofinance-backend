package util

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"fmt"
	"strings"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)


var cases = []testCase{/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func func TestGetTokenInHeaderAndVerify(t *testing.T) {

	var tests = []struct {
		name         string
		token        string
		headerFormat string
		want         error
	}{
		{"Successful Token Verification", "valid.jwt.token", "Bearer {token}", nil},
		{"JWT Token Missing In Header", "", "", fmt.Errorf("index out of range")},
		{"Invalid JWT Token In Header", "invalid.jwt.token", "Bearer {token}", fmt.Errorf("invalid token")},
		{"Malformed Authorization Header", "malformed.jwt.token", "{token}", fmt.Errorf("invalid token")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			t.Log("Starting Test Case: ", tt.name)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/dummy/route", nil)

			req.Header.Set("Authorization", strings.ReplaceAll(tt.headerFormat, "{token}", tt.token))

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			err := GetTokenInHeaderAndVerify(c)

			if tt.want != nil {
				if err == nil {
					t.Log("FAIL - Error expected was not received")
					t.Fail()
				}
				assert.Equal(t, tt.want.Error(), err.Error())
			} else if err != nil {
				t.Log("FAIL - Error was not expected but it was received")
				t.Fail()
			}

			t.Log("PASS")
		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func func TestValidateToken(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)
			err := ValidateToken(c, tc.token)
			if err == nil && tc.expectErr != nil {
				t.Fatalf("Expected error but got nil")
			} else if err != nil && tc.expectErr == nil {
				t.Fatalf("Expected no error but got %v", err)
			} else if err != nil && tc.expectErr != nil && !errors.Is(err, tc.expectErr) {
				t.Fatalf("Expected error %v but got %v", tc.expectErr, err)
			}
			if rr.Code != tc.statusCode {
				t.Fatalf("Expected status code %d but got %d", tc.statusCode, rr.Code)
			}
			t.Logf("%s passed", tc.name)
		})
	}
}

