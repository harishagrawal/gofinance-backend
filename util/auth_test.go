/*
ROOST_METHOD_HASH=GetTokenInHeaderAndVerify_c6fc249681
ROOST_METHOD_SIG_HASH=GetTokenInHeaderAndVerify_4459fbc010
*/
func TestGetTokenInHeaderAndVerify(t *testing.T) {
  testCases := []struct {
    name          string
    token         string
    errExpected   bool
    errMessage    string
  }{
    {
      name:          "Valid Authorization Header and Token",
      token:         "Bearer valid-token",
      errExpected:   false,
      errMessage:    "",
    },
    {
      name:          "Missing Authorization Header",
      token:         "",
      errExpected:   true,
      errMessage:    "Authorization header is missing",
    },
    {
      name:          "Invalid Authorization Token",
      token:         "Bearer invalid-token",
      errExpected:   true,
      errMessage:    "Token is invalid",
    },
    {
      name:          "Malformatted Authorization Header",
      token:         "malformatted token",
      errExpected:   true,
      errMessage:    "malformatted authorization header",
    },
  }

  for _, test := range testCases {
    t.Run(test.name, func(t *testing.T) {
      recorder := httptest.NewRecorder()
      context, _ := gin.CreateTestContext(recorder)
      request, _ := http.NewRequest(http.MethodGet, "/", nil)

      if test.token != "" {
        request.Header.Set("Authorization", test.token)
      }

      context.Request = request
      err := GetTokenInHeaderAndVerify(context)

      if test.errExpected {
        assert.Error(t, err)
        assert.Equal(t, test.errMessage, err.Error())
      } else {
        assert.NoError(t, err)
      }
    })
  }
}
