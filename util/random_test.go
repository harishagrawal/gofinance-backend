package util

import (
	"testing"
	"strings"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	testcases := []struct {
		name   string
		length int
		want   bool
	}{
		{"Normal Operation - Random Email Generation", 5, true},
		{"Edge Case - Zero Length Email", 0, true},
		{"Edge Case - Negative Length", -1, false},
		{"Large Number Input", 10000, true},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tt.name)

			got := RandomEmail(tt.length)
			gotSlice := strings.Split(got, "@")

			if len(gotSlice[0]) != tt.length && tt.want {
				t.Fatalf("RandomEmail() = %v, want length %v", len(gotSlice[0]), tt.length)
			}

			if gotSlice[1] != "email.com" && tt.want {
				t.Fatalf("RandomEmail() = %v, want %v", gotSlice[1], "email.com")
			}

			if tt.length < 0 && tt.want {
				t.Fatalf("RandomEmail() failed to handle negative length")
			}
			t.Logf("Test case: %s passed", tt.name)
		})
	}
}

