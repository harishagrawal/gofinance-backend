package util

import (
	"fmt"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"time"
)

var jwtSignedKey = []byte("secret_key")
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
type Context struct {
	writermem	responseWriter
	Request		*http.Request
	Writer		ResponseWriter
	Params		Params
	handlers	HandlersChain
	index		int8
	fullPath	string
	engine		*Engine
	params		*Params
	skippedNodes	*[]skippedNode
	mu		sync.RWMutex
	Keys		map // This mutex protects Keys map.
	// Keys is a key/value pair exclusively for the context of each request.
	[string]any
	Errors		errorMsgs
	Accepted	[ // Errors is a list of errors attached to all the handlers/middlewares who used this context.
	// Accepted defines a list of manually accepted formats for content negotiation.
	]string
	queryCache	url.Values
	formCache	url.Values
	sameSite	http.SameSite
}// queryCache caches the query result from c.Request.URL.Query().
// SameSite allows a server to define a cookie attribute making it impossible for
// the browser to send this cookie along with cross-site requests.


/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010


 */
func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)
	tokenToValidate := fields[1]
	errOnValiteToken := ValidateToken(ctx, tokenToValidate)
	if errOnValiteToken != nil {
		return errOnValiteToken
	}
	return nil
}

func TestGetTokenInHeaderAndVerify(t *testing.T) {
	dummyCorrectToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{"testuser", jwt.RegisteredClaims{}})
	dummyCorrectTokenString, err := dummyCorrectToken.SignedString(jwtSignedKey)
	if err != nil {
		t.Fatal("unable generate token for the dummy user")
	}

	tableTests := []struct {
		name          string
		tokenInHeader string
		expectFuncErr bool
	}{
		{"Valid Token Test", fmt.Sprintf("Bearer %s", dummyCorrectTokenString), false},
		{"Invalid Token Test", "Bearer wrongtoken", true},
		{"Authorization Header Missing", "", true},
		{"Empty AUTH token test", "Bearer ", true},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request.Header.Set("Authorization", tt.tokenInHeader)
			err := GetTokenInHeaderAndVerify(c)
			if tt.expectFuncErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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

/*
ROOST_METHOD_HASH=ValidateToken_7440899dfa
ROOST_METHOD_SIG_HASH=ValidateToken_ff3cc8ee02


 */
func TestValidateToken(t *testing.T) {

	var jwtSignedKey = []byte("secret_key")

	tests := []struct {
		name               string
		token              string
		expectErrorMessage string
	}{
		{
			"Valid Token Test",
			generateTestToken(jwtSignedKey, 30*time.Second),
			"",
		},
		{
			"Invalid Token Signature Test",
			generateTestToken([]byte("wrong_key"), 30*time.Second),
			jwt.ErrSignatureInvalid.Error(),
		},
		{
			"Invalid Token Format Test",
			"badToken",
			"token contains an invalid number of segments",
		},
		{
			"Expired Token Test",
			generateTestToken(jwtSignedKey, -30*time.Second),
			"Token is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(w)

			err := ValidateToken(ctx, tt.token)

			if tt.expectErrorMessage != "" {
				if err == nil || err.Error() != tt.expectErrorMessage {
					t.Errorf("Expected error message to be: '%s', got: '%v'", tt.expectErrorMessage, err)
				}
			} else if w.Code != http.StatusOK {
				t.Errorf("Expected status code to be: %d, got: %d", http.StatusOK, w.Code)
			}
		})
	}
}

func generateTestToken(jwtSignedKey []byte, expiresIn time.Duration) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: "TestUser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(expiresIn)),
		},
	})

	tokenString, _ := token.SignedString(jwtSignedKey)
	return tokenString
}

