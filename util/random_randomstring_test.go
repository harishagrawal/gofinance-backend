// ********RoostGPT********
/*
Test generated by RoostGPT for test go-calc using AI Type Azure Open AI and AI Model roostgpt-4-32k

ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

Scenario 1: RandomString produces a string of the correct length

Details: 
   Description: This test checks if the returned string is of correct length given an input parameter, `number`.
Execution:
   Arrange: We need to define a `number` integer.
   Act: Call `RandomString(number)` function.
   Assert: Check if the length of the returned string matches the `number` parameter.
Validation:
   For this test, the length of the output string is compared to the input `number` parameter using the `len()` function from the Go standard library. We expect them to match since the function represents generating a string of arbitrary characters of length `number`. This test is important to make sure the function behaves as expected under normal conditions.

Scenario 2: RandomString produces a string with only the specified characters

Details: 
   Description: This test will check that the string returned by `RandomString` contains only characters from the `alphabet` variable.
Execution:
   Arrange: Define a `number` integer. Make note of the `alphabet` variable.
   Act: Call `RandomString(number)` function.
   Assert: Check if all characters in the string are within `alphabet`.
Validation:
   In this test, the output string is iterated and each character is checked to see if it exists in the `alphabet` string. We expect all characters to be in `alphabet` as per the functionality of the function. This test is necessary for validating that only the specific characters included in the `alphabet` are being used.

Scenario 3: RandomString produces different strings for different calls

Details: 
   Description: This test is to check the randomness of the generated strings.
Execution:
   Arrange: A `number` integer needs to be defined.
   Act: Call `RandomString(number)` function multiple times.
   Assert: Each call should return a different string.
Validation:
   This test scenario checks if the function returns different strings for different method calls. The premise behind this test is that the `rand.Intn(k)` function should return a pseudorandom number each time it's called, leading to different strings being generated for different method calls. This test scenario is important to check the randomness of the generated string.

Scenario 4: RandomString handles a zero-length string parameter

Details: 
   Description: This test is to check how the function handles a zero-length string.
Execution:
   Arrange: Set the `number` parameter to 0.
   Act: Call `RandomString(number)` function.
   Assert: The return string should be an empty string.
Validation:
   In this scenario, we are testing if the function correctly handles the edge case when the `number` parameter is `0`. The function should return an empty string in this case. This test scenario is important for testing how the function handles edge cases.

Scenario 5: RandomString handles a negative integer parameter

Details: 
   Description: This test is to check how the function handles a negative `number` parameter.
Execution:
   Arrange: Set the `number` parameter to a negative integer.
   Act: Call `RandomString(number)` function.
   Assert: The return string should be an empty string.
Validation:
   This test validates the function's behavior for negative inputs. If the `number` parameter is negative, the function is expected to return an empty string as it signifies that no character string is requested. This test scenario is crucial since it is checking for a valid edge case and solidifies the function's robustness.
*/

// ********RoostGPT********
package util

import (
	"testing"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Definition of function being tested
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().Unix())

	tests := []struct {
		name         string
		input        int
		expectEmpty  bool
		isInAlphabet func(rune) bool
	}{
		{
			name:         "RandomString produces a string of the correct length",
			input:        5,
			expectEmpty:  false,
			isInAlphabet: func(r rune) bool { return strings.ContainsRune(alphabet, r) },
		},
		{
			name:         "RandomString produces a string with only specific characters",
			input:        7,
			expectEmpty:  false,
			isInAlphabet: func(r rune) bool { return strings.ContainsRune(alphabet, r) },
		},
		{
			name:         "RandomString produces different strings for different calls",
			input:        8,
			expectEmpty:  false,
			isInAlphabet: func(r rune) bool { return strings.ContainsRune(alphabet, r) },
		},
		{
			name:         "RandomString handles a zero-length string parameter",
			input:        0,
			expectEmpty:  true,
			isInAlphabet: func(r rune) bool { return strings.ContainsRune(alphabet, r) },
		},
		{
			name:         "RandomString handles a negative integer parameter",
			input:        -5,
			expectEmpty:  true,
			isInAlphabet: func(r rune) bool { return strings.ContainsRune(alphabet, r) },
		},
	}

	previousString := ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.input)

			if tt.expectEmpty && len(result) != 0 {
				t.Errorf("expected an empty string but got %s", result)
			}

			if !tt.expectEmpty && (len(result) != tt.input || result == previousString) {
				t.Errorf("string length mismatch or not producing different string for different calls. Got: %s but wanted length: %d and not equals to previous string: %s", result, tt.input, previousString)
			}

			for _, c := range result {
				if !tt.isInAlphabet(c) {
					t.Errorf("unexpected character: %c", c)
				}
			}

			previousString = result
		})
	}
}
