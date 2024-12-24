package util

import (
	"testing"
	"regexp"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)


var emailRegex = regexp.MustCompile(`^[a-z]*@email\.com$`)
var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	type test struct {
		name  string
		input int
		want  string
	}

	tests := []test{
		{name: "Valid Random Email Generation", input: 5},
		{name: "Edge Case of Zero Length Email Prefix", input: 0, want: "@email.com"},
		{name: "Negative Number Input", input: -3},
		{name: "Large Number Input", input: 1000000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got := RandomEmail(tc.input)

			match := emailRegex.MatchString(got)

			if tc.input < 1 {
				if got != tc.want {
					t.Errorf("Got %v, want %v", got, tc.want)
				}
			} else {
				if !match {
					t.Errorf("Generated Email %v did not match the expected format", got)
				}
			}

			if len(got) > tc.input+10 {
				t.Errorf("Email length %v exceeded the expected length %v", len(got), tc.input+10)
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

	type scenario struct {
		name            string
		input           int
		expectPanic     bool
	}
	
	scenarios := []scenario{
		{
			"Generate random string with a specified length",
			5,
			false,
		},
		{
			"Generate a string with zero length",
			0,
			false,
		},
		{
			"Generate a very long random string",
			10000,
			false,
		},
		{
			"Generate random string with a negative length",
			-5,
			true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			if s.expectPanic {
				assert.Panics(t, func() { RandomString(s.input) }, "The code did not panic")
			} else {
				res := RandomString(s.input)
				if len(res) != s.input {
					t.Errorf("Unexpected result length. Expected %d, but got %d.", s.input, len(res))
				}

			
				if strings.IndexFunc(res, func(c rune) bool { return !strings.ContainsRune(alphabet, c) }) != -1 {
					t.Errorf("Generated string contains unexpected characters.")
				}
			}
		})
	}
}

