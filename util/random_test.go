package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"github.com/wil-ckaew/gofinance-backend/util"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    int
		expected func(input string) bool
	}{
		{
			name:  "Generate an email with random string length",
			input: 5,
			expected: func(input string) bool {
				return len(strings.Split(input, "@")[0]) == 5
			},
		},
		{
			name:  "Generate an email with zero string length",
			input: 0,
			expected: func(input string) bool {
				return input == "@email.com"
			},
		},
		{
			name:  "Negative number input",
			input: -5,
			expected: func(input string) bool {
				return input == "@email.com"
			},
		},
		{
			name:  "Maximum number input",
			input: int(^uint(0)>>1) - 10,
			expected: func(input string) bool {
				return len(input) == int(^uint(0)>>1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := util.RandomEmail(test.input)
			if !test.expected(res) {
				t.Errorf("Failed to validate: %s", test.name)
			} else {
				t.Logf("Success: %s", test.name)
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
		{"Regular Functionality", 5, 5},
		{"Zero as Input", 0, 0},
		{"Negative Number as Input", -2, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			res := RandomString(tc.input)
			if len(res) != tc.expected {
				t.Fatalf("Unexpected result: got %v, want %v", len(res), tc.expected)
			}

			t.Log("Test case passed for ", tc.name)
		})
	}
}

