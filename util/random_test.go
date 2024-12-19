package util

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name       string
		number     int
		wantPrefix int
	}{
		{"Valid Input Test", 5, 5},
		{"Zero Number of Characters Test", 0, 0},
		{"Negative Number Test", -1, 0},
		{"Stress Test with Large Input Value", 1000000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Executing:", tt.name)

			got := RandomEmail(tt.number)
			gotBeforeAt := strings.Split(got, "@")[0]
			if len(gotBeforeAt) != tt.wantPrefix {
				t.Errorf("RandomEmail(%v) = %v, want prefix of length %v", tt.number, got, tt.wantPrefix)
			} else {
				t.Logf("Success: Generated email with prefix of expected length %v", tt.wantPrefix)
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

