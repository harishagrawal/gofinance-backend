package util

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
	"strings"
)









/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

FUNCTION_DEF=func RandomString(number int) string 

 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{0, 0},
		{10000, 10000},
		{-5, 0},
	}

	for _, test := range testCases {

		result := RandomString(test.input)

		if len(result) != test.expected {
			t.Errorf("RandomString(%d) failed; expected %d but got %d", test.input, test.expected, len(result))
		} else {
			t.Logf("Successfully got expected length of string for: RandomString(%d)", test.input)
		}

		for _, char := range result {
			if !unicode.IsLetter(char) {
				t.Errorf("RandomString(%d) failed; expected only alphabets but got non-alphabet character", test.input)
			}
		}

		t.Logf("RandomString(%d) passed ", test.input)
	}
}


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

FUNCTION_DEF=func RandomEmail(number int) string 

 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	testCases := []struct {
		detail, name string
		length       int
		expectedMail string
		expectError  bool
	}{
		{"Normal Operation with Positive Integer", "positiveInt", 7, "", false},
		{"Edge Case with Zero", "zeroInt", 0, "@email.com", false},
		{"Error Handling with Negative Integer", "negativeInt", -7, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.detail)

			email := RandomEmail(tc.length)

			if !tc.expectError {
				if tc.length > 0 && (len(strings.Split(email, "@")[0]) != tc.length) {
					t.Errorf("expected an email of length '%d', got '%s'", tc.length, email)
				}
				if tc.length == 0 && email != tc.expectedMail {
					t.Errorf("expected an email of value '%s', got '%s'", tc.expectedMail, email)
				}
			} else {
				if email != "" {
					t.Errorf("expected an invalid operation, but did not get one")
				}
			}
		})
	}

	t.Log("Randomness Test")
	email1 := RandomEmail(7)
	email2 := RandomEmail(7)

	if email1 == email2 {
		t.Error("expected two randomly generated strings to be distinct, but got two identical strings")
	}
}

