package util_test

import (
	"testing"
	"strings"
	"util"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name     string
		number   int
		expected int
	}{
		{name: "Generate random email with positive length", number: 10, expected: 10},
		{name: "Generate random email with zero length", number: 0, expected: 0},
		{name: "Generate random email with large positive length", number: 1000, expected: 1000},
		{name: "Generate random email with negative length", number: -10, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.number)

			parts := strings.Split(email, "@")
			got := len(parts[0])
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}

			if !strings.Contains(email, "@") {
				t.Errorf("expected an email address, got %s", email)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	tests := []struct {
		name     string
		length   int
		expected int
		err      error
	}{
		{"Random String Generation Test", 5, 5, nil},
		{"Zero Length String Test", 0, 0, nil},
		{"Negative Length String Test", -1, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := util.RandomString(tt.length)

			if len(res) != tt.expected {
				t.Errorf("Expected string of length %d but got %d", tt.expected, len(res))
			}

			for _, char := range res {
				if !strings.ContainsRune(util.alphabet, rune(char)) {
					t.Errorf("Random string contains an invalid character: %v", string(char))
				}
			}
			t.Log("Test Scenario Passed:", tt.name)
		})
	}
}

