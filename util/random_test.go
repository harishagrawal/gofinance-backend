package util

import (
	"testing"
	"strings"
	"math/rand"
	"time"
	"fmt"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	tests := []struct {
		name     string
		number   int
		expected string
	}{
		{"Normal operation with valid integer argument", 3, "@email.com"},
		{"Edge case with zero as argument", 0, "@email.com"},
		{"Edge case with negative integer as argument", -3, "@invalid"},
		{"Check for truly random behavior of the function", 3, "@email.com"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := RandomEmail(test.number)

			switch {

			case strings.Contains(got, "@email.com") && len(got) == (10+test.number):
				t.Logf("PASSED: The returned email address %v is in the expected format and contains a random string of %v characters", got, test.number)

			case test.number == 0 && got == "@email.com":
				t.Logf("PASSED: The returned email address is %v when input is zero", got)

			case test.number < 0 && got == "@invalid":
				t.Logf("PASSED: The returned email address is %v when input is negative", got)

			default:
				t.Errorf("FAILED: The returned email address %v is not as expected. Expected: %v", got, test.expected)
			}

			if test.name == "Check for truly random behavior of the function" {
				got2 := RandomEmail(test.number)
				if got == got2 {
					t.Errorf("FAILED: The function did not generate a truly random email address. %v is equal to %v", got, got2)
				}
				t.Logf("PASSED: The function successfully generated truly random email addresses")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	scenarios := []struct {
		desc     string
		length   int
		expected int
	}{
		{"Random String is Constructed with Desired Length", 10, 10},
		{"Random String Comprised Only of Alphabet Characters", 20, 20},
		{"Random String Generated is Unique on Subsequent Calls", 30, 30},
		{"Random String Generation with Zero Length", 0, 0},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			result := RandomString(scenario.length)

			if len(result) != scenario.expected {
				t.Errorf("Failed scenario '%s': Expected length %v but got %v", scenario.desc, scenario.expected, len(result))
			}

			for _, char := range result {
				if !strings.Contains(alphabet, string(char)) {
					t.Errorf("Failed scenario '%s': Non alphabet character %v introduced into random string", scenario.desc, string(char))
				}
			}

			secondResult := RandomString(scenario.length)
			if result == secondResult {
				t.Errorf("Failed scenario '%s': Non unique strings generated '%s' and '%s'", scenario.desc, result, secondResult)
			}

			if scenario.length == 0 && result != "" {
				t.Errorf("Failed scenario '%s': Expected empty string for zero length got %v", scenario.desc, result)
			}

			t.Log("Passed ", scenario.desc)
		})
	}
}

