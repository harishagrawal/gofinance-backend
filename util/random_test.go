package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name           string
		number         int
		expectedLength int
		expectedError  bool
	}{
		{
			name:           "Valid input test",
			number:         5,
			expectedLength: 5,
		},
		{
			name:           "Edge Case Test with zero number of characters input",
			number:         0,
			expectedLength: 0,
		},
		{
			name:          "Edge Case Test with negative input for the number of characters",
			number:        -5,
			expectedError: true,
		},
		{
			name:           "Stress Test with very large number of characters",
			number:         1000000,
			expectedLength: 1000000,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Executing scenario:", tc.name)

			result := RandomEmail(tc.number)

			if tc.expectedError && result != "" {
				t.Errorf("Unexpected result for input %d, got: %s, want: ''", tc.number, result)
			} else if !tc.expectedError && len(strings.Split(result, "@")[0]) != tc.expectedLength {
				t.Errorf("Unexpected length of email for input %d, got: %d, want: %d", tc.number, len(strings.Split(result, "@")[0]), tc.expectedLength)
			} else {
				t.Log("Test succeeded")
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
		description string
		input       int
		expectEmpty bool
	}{
		{
			description: "Testing RandomString with a valid number",
			input:       5,
			expectEmpty: false,
		},
		{
			description: "Testing RandomString with zero as input",
			input:       0,
			expectEmpty: true,
		},
		{
			description: "Testing RandomString with a negative number as input",
			input:       -5,
			expectEmpty: true,
		},
		{
			description: "Testing RandomString with a large number as input",
			input:       100000,
			expectEmpty: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := RandomString(tc.input)
			if tc.expectEmpty {
				assert.Empty(t, result, "Expected empty string but got string of length %d", len(result))
			} else {
				assert.Equal(t, tc.input, len(result), "Expected output of length %d but got %d", tc.input, len(result))
			}
		})
	}
}

