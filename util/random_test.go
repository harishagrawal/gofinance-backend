package util

import (
	rand "math/rand"
	strings "strings"
	testing "testing"
	time "time"
)

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

FUNCTION_DEF=func RandomString(number int) string
*/
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{
			name:     "Normal operation",
			length:   5,
			expected: 5,
		},
		{
			name:     "Zero length input",
			length:   0,
			expected: 0,
		},
		{
			name:     "Negative input",
			length:   -5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			got := RandomString(tt.length)
			if len(got) != tt.expected {
				t.Errorf("RandomString() = %v, want %v", len(got), tt.expected)
			} else {
				t.Logf("Test passed with %v", tt.name)
			}
		})
	}

	t.Run("Check randomness of output", func(t *testing.T) {
		def := 10
		got1 := RandomString(def)
		got2 := RandomString(def)

		if got1 == got2 {
			t.Errorf("RandomString() generated same string for two calls, randomness not ensured")
		} else {
			t.Log("Test passed for check randomness of output")
		}
	})
}

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

FUNCTION_DEF=func RandomEmail(number int) string
*/
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name          string
		number        int
		expectedError bool
	}{
		{
			name:          "Normal scenario",
			number:        5,
			expectedError: false,
		},
		{
			name:          "Edge case with zero",
			number:        0,
			expectedError: false,
		},
		{
			name:          "Negative number as argument",
			number:        -5,
			expectedError: true,
		},
		{
			name:          "Very large number as argument",
			number:        1000000,
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test. %v", r)
					t.Fail()
				}
			}()

			result := RandomEmail(tc.number)

			if tc.expectedError {
				if len(result) != 10 {
					t.Fatalf("Unexpected output: expected length is 10, got length: %d", len(result))
				} else {
					t.Logf("PASSED: Output length is as expected: %d", len(result))
				}
			} else {
				if len(result) != tc.number+10 {
					t.Fatalf("Unexpected output: expected length is %d, got length: %d", tc.number+10, len(result))
				} else {
					t.Logf("PASSED: Output length is as expected: %d", len(result))
				}
			}

			if !strings.Contains(result, "@email.com") {
				t.Fatalf("Unexpected output: expected string contains '@email.com', got: %s", result)
			} else {
				t.Logf("PASSED: Output contains '@email.com': %s", result)
			}
		})
	}
}
