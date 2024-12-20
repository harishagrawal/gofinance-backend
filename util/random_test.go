package util

import (
	"github.com/stretchr/testify/require"
	"testing"
	"strings"
	"math/rand"
	"time"
)

var sb strings.Builder/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expectedLen int
		expectError bool
	}{
		{
			name:        "Scenario 1: Normal Use Case for Random Email Generation",
			input:       10,
			expectedLen: 21,
			expectError: false,
		},
		{
			name:        "Scenario 2: Edge case when 0 is passed as a parameter",
			input:       0,
			expectedLen: 11,
			expectError: false,
		},
		{
			name:        "Scenario 3: Error handling when negative number is provided for function's parameter",
			input:       -4,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectError {
				require.Panics(t, func() { RandomEmail(tc.input) }, "Expected a panic for negative input")
				t.Log("Panic detected for negative test input as expected")
			} else {
				email := RandomEmail(tc.input)
				require.Contains(t, email, "@email.com", "email should contain '@email.com'")
				require.Len(t, email, tc.expectedLen, "email length should be input number plus length of '@email.com'")
				t.Logf("Email generated successfully with correct length")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name        string
		input       int
		expectedLen int
	}{{

		name:        "Zero characters",
		input:       0,
		expectedLen: 0,
	}, {

		name:        "Random string of length 3",
		input:       3,
		expectedLen: 3,
	}, {

		name:        "Negative number input",
		input:       -2,
		expectedLen: 0,
	}, {

		name:        "High number input",
		input:       1000000,
		expectedLen: 1000000,
	}}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running: %v", tc.name)

			got := RandomString(tc.input)

			if len(got) != tc.expectedLen {
				t.Errorf("got string length: %v but expected: %v", len(got), tc.expectedLen)
			} else {
				t.Logf("Validation successful - got string length: %v as expected", len(got))
			}
		})
	}
}

