package util

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	type testCase struct {
		name           string
		input          int
		expectedLength int
		expectedSuffix string
	}

	testCases := []testCase{
		{
			name:           "Scenario 1: Valid Random Email Generation",
			input:          5,
			expectedLength: 5 + len("@email.com"),
			expectedSuffix: "@email.com",
		},
		{
			name:           "Scenario 2: Invalid Random Email Generation for Negative Input",
			input:          -3,
			expectedLength: len("@email.com"),
			expectedSuffix: "@email.com",
		},
		{
			name:           "Scenario 3: Random Email Generation with Zero Length",
			input:          0,
			expectedLength: len("@email.com"),
			expectedSuffix: "@email.com",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			email := RandomEmail(test.input)

			assert.Equal(t, test.expectedLength, len(email), fmt.Sprintf("Expected to be %d but got %d", test.expectedLength, len(email)))

			assert.Equal(t, test.expectedSuffix, email[len(email)-len("@email.com"):], fmt.Sprintf("Expected to be @email.com but got %s", email[len(email)-len("@email.com"):]))
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	t.Run("positive length", func(t *testing.T) {
		t.Log("Scenario 1: Generate random string with positive length")
		const numberOfCharacters = 5
		str := RandomString(numberOfCharacters)

		if len(str) != numberOfCharacters {
			t.Errorf("RandomString(%v) = %v, expected length %v, got %v", numberOfCharacters, str, numberOfCharacters, len(str))
		}

		t.Log("Passed scenario 1")
	})

	t.Run("zero-length string", func(t *testing.T) {
		t.Log("Scenario 2: Edge Scenario - Generate a zero-length string")
		const numberOfCharacters = 0
		str := RandomString(numberOfCharacters)

		if len(str) != numberOfCharacters {
			t.Errorf("RandomString(%v) = %v, expected length %v, got %v", numberOfCharacters, str, numberOfCharacters, len(str))
		}

		t.Log("Passed scenario 2")
	})

	t.Run("large random string", func(t *testing.T) {
		t.Log("Scenario 3: Generate a large random string")
		const numberOfCharacters = 1000
		str := RandomString(numberOfCharacters)

		if len(str) != numberOfCharacters {
			t.Errorf("RandomString(%v) = %v, expected length %v, got %v", numberOfCharacters, str, numberOfCharacters, len(str))
		}

		t.Log("Passed scenario 3")
	})

	t.Run("randomness of generated strings", func(t *testing.T) {
		t.Log("Scenario 4: Randomness of the generated strings")
		const numberOfCharacters = 10
		const numberOfSamples = 10000
		set := make(map[string]bool, numberOfSamples)
		for i := 0; i < numberOfSamples; i++ {
			str := RandomString(numberOfCharacters)
			if _, exists := set[str]; exists {
				t.Errorf("Same string generated more than once: %v", str)
			}
			set[str] = true
		}

		t.Log("Passed scenario 4")
	})
}

