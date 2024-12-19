package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wil-ckaew/gofinance-backend/util"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang-jwt/jwt/v4"
)


var jwtTestCases = []struct {/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func func TestGetTokenInHeaderAndVerify(t *testing.T) {
	tt := []struct {
		name          string
		authorization string
		shouldErr     bool
	}{
		{
			name:          "Legitimate Token",
			authorization: "Bearer legitimate_test_token",
			shouldErr:     false,
		},
		{
			name:          "Token missing in Header",
			authorization: "",
			shouldErr:     true,
		},
		{
			name:          "Invalid Token in Header",
			authorization: "Bearer invalid_token",
			shouldErr:     true,
		},
		{
			name:          "Malformed Header",
			authorization: "BearerMalformed",
			shouldErr:     true,
		},
		{
			name:          "Check with empty Header",
			authorization: "",
			shouldErr:     true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.GET("/", func(c *gin.Context) {
				err := util.GetTokenInHeaderAndVerify(c)
				assert.Equal(t, tc.shouldErr, err != nil)
			})

			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("authorization", tc.authorization)

			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if tc.shouldErr {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			} else {
				assert.Equal(t, http.StatusOK, resp.Code)
			}

		})
	}
}

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func func TestValidateToken(t *testing.T) {

	for _, tt := range jwtTestCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			db, mock, _ := sqlmock.New()
			defer db.Close()

			if tt.setupClaims != nil || tt.setupMocks != nil {
				claims := &TestClaims{}
				if tt.setupClaims != nil {
					tt.setupClaims(claims)
				}
				if tt.setupMocks != nil {
					tt.setupMocks(mock)
				}
			}

			ValidateToken(c, tt.token)

			assert.Equal(t, tt.valid, c.GetHeader("token_valid") == "true")
			assert.Equal(t, tt.outputStatus, w.Code)
		})
	}
}

