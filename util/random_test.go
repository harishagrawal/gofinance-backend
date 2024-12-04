package util

import (
	"strings"
	"testing"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	scenarios := []struct {
		desc     string
		length   int
		expected bool
	}{
		{
			desc:     "Validating the length of the RandomEmail",
			length:   15,
			expected: true,
		},
		{
			desc:     "Ensuring string ends with @email.com",
			length:   10,
			expected: true,
		},
		{
			desc:     "Generating unique email addresses",
			length:   5,
			expected: true,
		},
		{
			desc:     "No email generated for negative input",
			length:   -1,
			expected: true,
		},
		{
			desc:     "Testing Scalability",
			length:   10000,
			expected: false,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.desc, func(t *testing.T) {
			email := RandomEmail(tt.length)

			switch tt.desc {
			case "Validating the length of the RandomEmail":
				if len(email) != tt.length+len("@email.com") {
					t.Errorf("Expected email length %d, got %d", tt.length+len("@email.com"), len(email))
				}

			case "Ensuring string ends with @email.com":
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("Generated email does not end with \"@email.com\"")
				}

			case "Generating unique email addresses":
				email2 := RandomEmail(tt.length)
				if email == email2 {
					t.Errorf("Generated emails are not unique")
				}

			case "No email generated for negative input":
				if len(email) != len("@email.com") {
					t.Errorf("Expected an empty string due to negative input, email generated")
				}

			case "Testing Scalability":
				if len(email) < tt.length {
					t.Errorf("Testing for a very large input failed. Email length less than expected")
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	if number < 0 {
		panic("length cannot be negative")
	}

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "Random String of specific length",
			length: 10,
		},
		{
			name:   "Empty Random String",
			length: 0,
		},
		{
			name:   "Negative length input",
			length: -5,
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if tt.length >= 0 {
						t.Errorf("The code paniced with valid length input")
					}
				} else if tt.length < 0 {
					t.Errorf("The code did not panic with negative length input")
				}
			}()
			result := RandomString(tt.length)

			if tt.length == 0 && result != "" {
				t.Errorf("got %s, want ''", result)
			}

			if len(result) != tt.length {
				t.Errorf("String length = %d; want %d", len(result), tt.length)
			}

			for _, char := range result {
				if !strings.Contains(alphabet, string(char)) {
					t.Errorf("String contains invalid character: %v", char)
				}
			}
		})
	}
}

