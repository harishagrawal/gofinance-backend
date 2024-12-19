package util

import (
	"strings"
	"testing"
	"regexp"
	"fmt"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func func TestRandomEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       int
		expectedLen int
	}{
		{
			name:        "Normal operation with valid input",
			input:       5,
			expectedLen: 5,
		},
		{
			name:        "Edge case with zero input",
			input:       0,
			expectedLen: 0,
		},
		{
			name:        "Edge case with negative input",
			input:       -3,
			expectedLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)
			emailParts := strings.Split(email, "@")
			localPart := emailParts[0]

			if len(localPart) != tt.expectedLen {
				t.Errorf("RandomEmail(%v) = %v, want length = %v", tt.input, email, tt.expectedLen)
				t.Logf("Test %s failed as the length of email's local part did not match the input number", tt.name)
			} else {
				t.Logf("Test %s passed", tt.name)
			}

			matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
			if err != nil {
				t.Logf("Error matching regex: %s", err)
			}

			if !matched {
				t.Errorf("RandomEmail(%v) = %v is not a valid email format", tt.input, email)
				t.Logf("Test %s failed as the output email did not match the valid email format", tt.name)
			} else {
				t.Logf("Test %s passed", tt.name)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func func TestRandomString(t *testing.T) {

	var tests = []struct {
		name     string
		input    int
		expected int
	}{
		{"Normal operation with positive non-zero number", 5, 5},
		{"Edge case operation with zero number", 0, 0},
		{"Edge case operation with negative number", -3, 0},
		{"Stress testing with a large number", 1_000_000, 1_000_000},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := RandomString(tt.input)

			assert.Equal(t, tt.expected, len(result), "They should be equal in length")

			if len(result) != tt.expected {
				t.Logf("FAILED: %s: Expected '%d' but got value '%d'", tt.name, tt.expected, len(result))
			} else {
				t.Logf("SUCCESS: %s: Expected '%d' and got '%d'", tt.name, tt.expected, len(result))
			}
		})
	}
}

