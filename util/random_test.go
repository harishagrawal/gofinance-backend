package util

import (
	"strings"
	"testing"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd
 */

func TestRandomEmail(t *testing.T) {

	t.Run("Testing length of generated email", func(t *testing.T) {
		length := 10
		email := RandomEmail(length)
		if len(email) != length+11 {
			t.Errorf("Expected length of %d but got %d", length+11, len(email))
		} else {
			t.Logf("Generated email is of correct length")
		}
	})

	t.Run("Testing domain of generated email", func(t *testing.T) {
		email := RandomEmail(10)
		if !strings.Contains(email, "@email.com") {
			t.Errorf("Expected domain '@email.com' not found in email")
		} else {
			t.Logf("Domain check is successful")
		}
	})

	t.Run("Testing randomness of emails", func(t *testing.T) {
		emails := make(map[string]bool)
		for i := 0; i < 10; i++ {
			email := RandomEmail(10)
			if emails[email] {
				t.Errorf("Email is repeated: %s", email)
			}
			emails[email] = true
		}
		t.Logf("Randomness check is successful")
	})

	t.Run("Testing negative number input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code should have panicked")
			} else {
				t.Logf("Handled negative input gracefully with a panic")
			}
		}()

		RandomEmail(-10)
	})
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a
 */

func TestRandomString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Valid input",
			input:    5,
			expected: 5,
		},
		{
			name:     "Zero input",
			input:    0,
			expected: 0,
		},
		{
			name:     "Negative input",
			input:    -1,
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.expected==-1 {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Code should have panicked")
					} else {
						t.Logf("Handled negative input gracefully with a panic")
					}
				}()
			}
			
			result := RandomString(tc.input)

			if len(result) != tc.expected && tc.expected!=-1 {
				t.Errorf("test '%s' failed: expected string of length '%d', got string of length '%d'", tc.name, tc.expected, len(result))
			} else {
				t.Logf("test '%s' passed: correct string length of '%d' was generated", tc.name, len(result))
			}
		})
	}
}
