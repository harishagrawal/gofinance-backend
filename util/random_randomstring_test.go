// ********RoostGPT********
/*
Test generated by RoostGPT for test go-calc using AI Type Azure Open AI and AI Model roostgpt-4-32k

ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

Scenario 1: Generating Random Strings

Details:
  Description: Test is meant to check the actual functionality of the function in its basic operation. It would confirm if the function correctly generate a string of random characters of the specified length. 
Execution:
  Arrange: No set up or mock data is required as the function does not depend on any external dependencies.
  Act: Invoke RandomString with a number parameter.
  Assert: Verify that the returned string is of the expected length.
Validation:
  This assertion checks if the random string generated is of the right length. It's essential to the application's behavior or business requirements as this confirms that the function behaves as expected in a normal operation.

Scenario 2: Generating Zero Length String

Details:
  Description: This test is meant to handle bounds limit of the function. It would confirm if the function handles the edge case of zero length correctly, i.e., it should return an empty string.
Execution:
  Arrange: No set up or mock data is required here as well.
  Act: Invoke RandomString with 0 as the parameter.
  Assert: Verify that the function returns an empty string.
Validation:
  This assertion checks if the function behaves correctly when provided with zero as an input. It's of importance as it ensures that the function has been designed and implemented to handle edge cases correctly.

Scenario 3: Generating Long String

Details:  
  Description: This test is to check whether the function can handle the generation of a long string without crashing.
Execution:
  Arrange: No set up or mock data required.
  Act: Invoke RandomString with a very large number as the parameter.
  Assert: Verify the function returns a string of the specified length.
Validation:
  This assertion verifies if the function can handle generating very large strings. It's important to ensure that the application has an understanding of the functions limitations in terms of resource usage when generating large strings.

Scenario 4: Negative Length String

Details:
  Description: This test is to verify the behavior of the function when it is provided with a negative length for the string. Ideally, it should either return an error or an empty string in this case.
Execution:
  Arrange: No set up or mock data required.
  Act: Invoke RandomString with a negative number as the parameter.
  Assert: Observe the function's behavior (depending on the implementation, you may check for an empty string, or if function panics, you might handle panic).
Validation:
  This test ensures if the function handles erroneous input correctly. As this function should only accept positive integers as valid input, it is important to check and validate the function's response when provided with a negative number.

*/

// ********RoostGPT********
package util

import (
	"math/rand"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // random number seed; do not remove

	// define test cases
	testCases := []struct {
		name     string
		number   int
		expected func(t *testing.T, res string)
	}{
		{
			"Generating Random Strings",
			10,
			func(t *testing.T, res string) {
				length := len(res)
				if length != 10 {
					t.Errorf("expected length of 10, got %d", length)
				}
			},
		},
		{
			"Generating Zero Length String",
			0,
			func(t *testing.T, res string) {
				length := len(res)
				if length != 0 {
					t.Errorf("expected length of 0, got %d", length)
				}
			},
		},
		{
			"Generating Long String",
			10000,
			func(t *testing.T, res string) {
				length := len(res)
				if length != 10000 {
					t.Errorf("expected length of 10000, got %d", length)
				}
			},
		},
		{
			"Negative Length String",
			-1,
			func(t *testing.T, res string) {
				length := len(res)
				if length != 0 {
					t.Errorf("expected length of 0 when using a negative integer, got %d", length)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.number)
			tc.expected(t, result) // run test and determine pass/fail
		})
	}
}

// RandomString generates a random string of the given length.
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
