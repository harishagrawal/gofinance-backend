package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var sb strings.Builder/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	type testCase struct {
		name          string
		input         int
		expectedError error
	}

	testCases := []testCase{
		{"Valid Random Email Generation", 5, nil},
		{"Zero-Length Random Email Generation", 0, nil},
		{"Negative Number Input Validation", -5, fmt.Errorf("Received negative input")},
		{"Large Number Input", 10000, nil},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.input)

			if tt.input < 0 {
				if !strings.Contains(email, "Received negative input") {
					t.Errorf("unexpected error: want 'Received negative input', but got %s", email)
				}
			} else {
				if len(email) != tt.input+10 {
					t.Errorf("unexpected email length: want '%d', got '%d'", tt.input+10, len(email))
				}
				if !strings.Contains(email, "@email.com") {
					t.Error("email does not contain '@email.com'")
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

	rand.Seed(1)

	testCases := []struct {
		name           string
		input          int
		expectedLength int
	}{
		{"Standard Functionality", 10, 10},
		{"Zero Input", 0, 0},
		{"Negative Input", -10, 0},
		{"Performance Test", 1000000, 1000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Executing test scenario:", tc.name)

			result := RandomString(tc.input)

			if len(result) != tc.expectedLength {
				t.Errorf("expected string length of %d, got %d", tc.expectedLength, len(result))
			} else {
				t.Logf("Successfully passed test scenario: '%s' with string length of %d.", tc.name, len(result))
			}
		})
	}
}

