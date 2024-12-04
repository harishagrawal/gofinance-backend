package util

import (
	"fmt"
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

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name          string
		emailLength   int
		expectedCount int
	}{
		{"Random Email Length Validation", 5, 15},
		{"Random Email Length with Edge Case - Zero", 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			email := RandomEmail(tt.emailLength)
			if len(email) != tt.expectedCount {
				t.Errorf("Failed %s: Expected length %v but got %v", tt.name, tt.expectedCount, len(email))
			} else {
				t.Logf("Success %s: Expected length %v and got %v", tt.name, tt.expectedCount, len(email))
			}

			if tt.emailLength != 0 && !strings.HasSuffix(email, "@email.com") {
				t.Errorf("Failed %s: Expected format string@email.com but got %v", tt.name, email)
			} else if tt.emailLength == 0 && email != "@email.com" {
				t.Errorf("Failed %s: Expected @email.com for zero length but got %v", tt.name, email)
			} else {
				t.Logf("Success %s: Email format is correct", tt.name)
			}

			email2 := RandomEmail(tt.emailLength)
			if email == email2 {
				t.Errorf("Failed %s: Expected different emails for the same length but got same %v", tt.name, email)
			} else {
				t.Logf("Success %s: Different emails generated for the same length", tt.name)
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

	var tests = []struct {
		name     string
		input    int
		isLength bool
		noEquals bool
	}{
		{"Generate empty string with zero length", 0, true, false},
		{"Generate random string of given length", 10, true, false},
		{"Generate different strings on different calls", 10, false, true},
	}

	for _, tt := range tests {

		t.Log(tt.name)

		result := RandomString(tt.input)
		if tt.isLength {
			length := len(result)
			if length != tt.input {
				t.Errorf("RandomString(%d): expected length %d, actual length %d", tt.input, tt.input, length)
			} else {
				t.Logf("Successful Test! RandomString(%d): generated string with correct length", tt.input)
			}
		}
		if tt.noEquals {
			result2 := RandomString(tt.input)
			if result == result2 {
				t.Errorf("RandomString(%d): RandomString should return a different result each call. Received two equal strings", tt.input)
			} else {
				t.Logf("Successful Test! RandomString(%d): generated two distinct strings", tt.input)
			}
		}

		for _, char := range result {
			if !strings.Contains(string(alphabet), string(char)) {
				t.Errorf("RandomString(%d): the character %s is not part of the specified alphabet", tt.input, string(char))
			}
		}
	}
}

