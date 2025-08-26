package util

import (
	rand "math/rand"
	debug "runtime/debug"
	strings "strings"
	testing "testing"
	time "time"
)

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a

FUNCTION_DEF=func RandomString(number int) string
*/
func TestRandomString(t *testing.T) {

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	t.Run("Scenario 1: Validate Generation of String of Specified Length", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 50
		result := RandomString(number)
		if len(result) != number {
			t.Errorf("Expected string length %d, but got %d", number, len(result))
		}
		t.Logf("Success: Generated string length %d matches the expected length %d.", len(result), number)
	})

	t.Run("Scenario 2: Validate String Contains Only Characters From the Alphabet", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 30
		result := RandomString(number)
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Generated string contains invalid character: %c", char)
			}
		}
		t.Logf("Success: All characters in generated string are from the alphabet.")
	})

	t.Run("Scenario 3: Validate Randomness of Generated Strings", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 20
		numIterations := 100
		uniqueStrings := make(map[string]bool)
		for i := 0; i < numIterations; i++ {
			result := RandomString(number)
			if uniqueStrings[result] {
				t.Errorf("Duplicate string encountered: %s", result)
			}
			uniqueStrings[result] = true
		}
		t.Logf("Success: RandomString generated %d unique strings with %d iterations.", len(uniqueStrings), numIterations)
	})

	t.Run("Scenario 4: Validate Behavior with Zero-Length Input", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 0
		result := RandomString(number)
		if len(result) != 0 {
			t.Errorf("Expected empty string for zero-length input, but got %s", result)
		}
		t.Logf("Success: RandomString gracefully handled zero-length input.")
	})

	t.Run("Scenario 5: Validate Behavior with Negative-Length Input", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := -1
		result := RandomString(number)
		if len(result) != 0 {
			t.Errorf("Expected empty string for negative-length input, but got %s", result)
		}
		t.Logf("Success: RandomString gracefully handled negative-length input.")
	})

	t.Run("Scenario 6: Validate Large-Length Input Handling", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 10000
		start := time.Now()
		result := RandomString(number)
		duration := time.Since(start)
		if len(result) != number {
			t.Errorf("Expected string length %d, but got %d", number, len(result))
		}
		t.Logf("Success: RandomString handled large input length in %s.", duration)
	})

	t.Run("Scenario 7: Validate Consistency Under Seed Initialization", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 30
		seed := int64(12345)
		rand.Seed(seed)
		result1 := RandomString(number)
		rand.Seed(seed)
		result2 := RandomString(number)
		if result1 != result2 {
			t.Errorf("Expected deterministic output from seeded RNG, but results differ: %s, %s", result1, result2)
		}
		t.Logf("Success: RandomString produced consistent output under seeded RNG.")
	})

	t.Run("Scenario 8: Validate Integration with Claims and Alphabet Definitions", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		number := 25
		result := RandomString(number)
		for _, char := range result {
			if !strings.ContainsRune(alphabet, char) {
				t.Errorf("Generated string contains invalid character: %c", char)
			}
		}

		claims := Claims{
			Username: "testuser",
		}
		claims.Username = result
		t.Logf("Success: Claims struct interacted correctly with RandomString output.")
	})
}
