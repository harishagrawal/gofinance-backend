package util

import (
	"strings"
	"testing"
	"os"
	"fmt"
	"github.com/stretchr/testify/assert"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	tests := []struct {
		name   string
		number int
		expect string
	}{
		{
			name:   "Valid Output Test Scenario",
			number: 5,
			expect: "@email.com",
		},
		{
			name:   "Edge Case - Zero Value Input",
			number: 0,
			expect: "@email.com",
		},
		{
			name:   "Edge Case - Negative Number Input",
			number: -5,
			expect: "@email.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RandomEmail(test.number)

			if strings.HasSuffix(got, test.expect) && len(got) == test.number+len(test.expect) {
				t.Logf("Test Passed. Expected suffix of '%s', got '%s'", test.expect, got)
			} else {
				t.Errorf("Test Failed. Expected suffix of '%s', got '%s'", test.expect, got)
			}
		})
	}

	t.Run("Non-deterministic Functionality", func(t *testing.T) {
		emailSet := make(map[string]struct{})
		number := 5
		for i := 0; i < 10; i++ {
			email := RandomEmail(number)
			if _, exists := emailSet[email]; exists {
				t.Errorf("Test Failed. Expected a unique email, got a duplicate '%s'", email)
			} else {
				emailSet[email] = struct{}{}
				t.Logf("Test Passed. Got a unique email '%s'", email)
			}
		}
	})
}


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	testCases := []struct {
		length int
		desc   string
	}{
		{5, "Random String Generations of Varying Lengths"},
		{10, "Random String Generations of Varying Lengths"},
		{50, "Random String Generations of Varying Lengths"},
		{0, "Checking Output for Edge Case of Zero Length"},
		{-5, "Negative Input Check"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			origStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			result := RandomString(tc.length)

			outC := make(chan string)
			go func() {
				var buf strings.Builder
				fmt.Fscanf(r, "%s", &buf)
				outC <- buf.String()
			}()
			w.Close()
			os.Stdout = origStdout
			out := <-outC

			if tc.length >= 0 {
				assert.Equal(t, tc.length, len(result), "Incorrect length for test '%s'", tc.desc)
				assert.Equal(t, tc.length, len(out), "Output length doesn't match input in test '%s'", tc.desc)
				assert.True(t, doesContain(result), "The string doesn't contain the alphabet '%s'", tc.desc)
			} else {
				t.Log("Negative Input doesn't throw an error. It should be handled accordingly.")
			}

			if tc.desc == "Random String Generations of Varying Lengths" {
				nextResult := RandomString(tc.length)
				assert.NotEqual(t, result, nextResult, "Repetitive Calls Return Same Strings")
			}
		})
	}
}

func doesContain(result string) bool {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for _, r := range result {
		if !strings.Contains(alphabet, string(r)) {
			return false
		}
	}
	return true
}

