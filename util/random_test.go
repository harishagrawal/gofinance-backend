package util

import (
	"math/rand"
	"testing"
	"time"
	"strings"
)









/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

FUNCTION_DEF=func RandomString(number int) string 

 */
func TestRandomString(t *testing.T) {
	tt := []struct {
		name     string
		number   int
		expected int
	}{
		{"Scenario 1: Testing with positive number", 5, 5},
		{"Scenario 2: Testing with zero", 0, 0},
		{"Scenario 3: Testing with negative number", -7, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := RandomString(tc.number)
			if x := len(res); x != tc.expected {
				t.Fatalf("RandomString(%d) = length of %v; want %v", tc.number, x, tc.expected)
			} else {
				t.Logf("Success: RandomString(%d) = length of %v as expected", tc.number, x)
			}
		})
	}

	t.Run("Scenario 4: Testing randomness", func(t *testing.T) {
		n := 5
		res1 := RandomString(n)
		res2 := RandomString(n)

		if res1 == res2 {
			t.Fatalf("Random string generator failed on randomness, generated same strings: %v", res1)
		} else {
			t.Logf("Success: The RandomString function produces random outputs.")
		}
	})
}


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

FUNCTION_DEF=func RandomEmail(number int) string 

 */
func TestRandomEmail(t *testing.T) {

	var testCases = []struct {
		description string
		input       int
		lenExpected int
	}{
		{
			description: "Scenario 1: Validate the return of a valid email string when a positive number is passed",
			input:       3,
			lenExpected: 20,
		},
		{
			description: "Scenario 2: Validate the return of a valid email string when zero is passed",
			input:       0,
			lenExpected: 10,
		},
		{
			description: "Scenario 3: Validate the return of a valid email string when a negative number is passed",
			input:       -3,
			lenExpected: 10,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			t.Log("Arrange and Act: Get random email from RandomEmail function.")

			email := RandomEmail(tt.input)
			emailName := strings.Split(email, "@")[0]

			t.Logf("Returned Email: %v", email)

			if got := len(emailName); got != tt.lenExpected {
				t.Errorf("Failed: %s: RandomEmail(%d): Expected email length before '@': %d, got: %d",
					tt.description, tt.input, tt.lenExpected, got)
			} else {
				t.Logf("Success: %s: RandomEmail(%d): Expected email length before '@': %d, got: %d",
					tt.description, tt.input, tt.lenExpected, got)
			}
		})
	}
}

