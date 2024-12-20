package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"regexp"
	"time"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	scenarios := []struct {
		name       string
		length     int
		shouldPass bool
	}{
		{"Normal operation with valid parameter", 5, true},
		{"Zero as parameter", 0, true},
		{"Parameter exceeding limit", 3000, false},
		{"Negative integer as parameter", -5, false},
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			t.Log("Executing scenario:", s.name)
			email := RandomEmail(s.length)

			if !s.shouldPass {
				if len(email) > 254 || len(email) <= 0 || !re.MatchString(email) {
					t.Log("Scenario PASSED as expected since input did not meet valid email requirements")
				} else {
					t.Errorf("Scenario FAILED. Expected invalid email but received: %s", email)
				}
			} else {

				if len(email) == len(fmt.Sprintf("%s@email.com", strings.Repeat("x", s.length))) && re.MatchString(email) {
					t.Log("Scenario PASSED as expected with properly formatted email generated")
				} else {
					t.Errorf("Scenario FAILED. Expected a valid email but received: %s", email)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "String Length Validation Test",
			input:    5,
			expected: 5,
		},
		{
			name:     "Zero Length String Test",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actual := RandomString(tc.input)

			if len(actual) != tc.expected {
				t.Errorf("Failed %v: Expected string length %v, Got %v", tc.name, tc.expected, len(actual))
			}

			for _, char := range actual {
				if !strings.Contains(alphabet, string(char)) {
					t.Errorf("Failed %v: All characters must be from alphabet", tc.name)
				}
			}
		})
	}
}

