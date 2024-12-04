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

	tests := []struct {
		name     string
		number   int
		expected int
	}{
		{"Scenario 1", 10, 10},
		{"Scenario 2", 0, 0},
		{"Scenario 3", 10000, 10000},
		{"Scenario 4", -10, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			email := RandomEmail(test.number)

			splitEmail := strings.Split(email, "@")
			actual := len(splitEmail[0])

			if actual != test.expected {
				t.Errorf("Expected random string to have a length of %d, but got a length of %d", test.expected, actual)
			}

			if !strings.Contains(email, "@") {
				t.Errorf("Expected random email to contain '@', but got %s", email)
			}

			if splitEmail[1] != "email.com" {
				t.Errorf("Expected random email to end with 'email.com', got %s", splitEmail[1])
			}

			t.Log("Success:", test.name)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	testTable := []struct {
		name  string
		input int
		match func(t *testing.T, result string)
	}{
		{
			name:  "Random String Generation Test",
			input: 10,
			match: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("RandomString(10) returned a string of length %d, want 10", len(result))
				}
			},
		},
		{
			name:  "Zero Length String Test",
			input: 0,
			match: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("RandomString(0) returned a string of length %d, want 0", len(result))
				}
			},
		},
		{
			name:  "Negative Length String Test",
			input: -10,
			match: func(t *testing.T, result string) {
				if len(result) != 0 {
					t.Errorf("RandomString(-10) returned a string of length %d, want 0", len(result))
				}
			},
		},
		{
			name:  "Random String Qualitative Test",
			input: 100,
			match: func(t *testing.T, result string) {
				for _, c := range result {
					if !strings.Contains(alphabet, string(c)) {
						t.Errorf("RandomString() returned a string with character %v, which is not in alphabet", c)
					}
				}
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.input)

			tc.match(t, result)
		})
	}
}

func init() {
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
}

