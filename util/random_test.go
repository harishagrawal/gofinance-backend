package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	tests := []struct {
		name     string
		argument int
	}{
		{
			"Valid number input for random email generation",
			5,
		},
		{
			"Zero input for random email generation",
			0,
		},
		{
			"Negative number input for random email generation",
			-5,
		},
		{
			"Large number input for random email generation",
			10000,
		},
	}

	for _, tt := range tests {
		email := RandomEmail(tt.argument)
		emailParts := strings.Split(email, "@")
		got := len(emailParts[0])

		switch {
		case tt.argument < 0 && got != 0:
			t.Errorf("Case: %s, Expected length: 0, Got length: %d", tt.name, got)
		case tt.argument >= 0 && got != tt.argument:
			t.Errorf("Case: %s, Expected length: %d, Got length: %d", tt.name, tt.argument, got)
		default:
			t.Logf("Test: %s Passed", tt.name)
		}

		t.Log(email)
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	var testCases = []struct {
		name      string
		number    int
		expectErr bool
	}{
		{"Normal Generation", 5, false},
		{"Empty String Generation", 0, false},
		{"Negative Integer as Argument", -5, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if !tc.expectErr {
						t.Errorf("RandomString() got unexpected panic: %v", r)
					}
				} else if tc.expectErr {
					t.Errorf("RandomString() did not panic but was expected to")
				}
			}()

			res := RandomString(tc.number)

			if len(res) != tc.number {
				t.Errorf("RandomString() got length = %d; want %d", len(res), tc.number)
			}

			if tc.number == 0 && res != "" {
				t.Errorf("RandomString() got = %s; want an empty string", res)
			}

			for _, c := range res {
				if !strings.Contains(alphabet, string(c)) {
					t.Errorf("RandomString() generated a string with a character outside the allowed set: got = %s; want in %s", string(c), alphabet)
				}
			}
		})
	}

	t.Run("Infinite Generation", func(t *testing.T) {
		t.Log("Infinite Generation test case requires manual testing due to possible infinite loop condition")
	})

	t.Run("Randomness of Generated Strings", func(t *testing.T) {
		s1 := RandomString(5)
		s2 := RandomString(5)

		if s1 == s2 {
			t.Errorf("RandomString() generated the same string twice: got s1 = %s; got s2 = %s", s1, s2)
		}
	})

}

