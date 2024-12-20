package util

import (
	"strings"
	"testing"
	"fmt"
	"math/rand"
	"time"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	tests := []struct {
		name       string
		input      int
		expectErr  bool
		expectSubs []string
	}{
		{"Normal Input", 5, false, []string{"@email.com"}},
		{"Zero Input", 0, false, []string{"@email.com"}},
		{"Negative Input", -1, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := RandomEmail(tt.input)
			t.Log("Running:", tt.name, "- input:", tt.input, "- expected error:", tt.expectErr)

			if tt.expectErr && err == nil {
				t.Errorf("expected an error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("did not expect an error but got %v", err)
			}
			if tt.expectErr {
				return
			}

			for _, subs := range tt.expectSubs {
				if !strings.Contains(email, subs) {
					t.Errorf("Expected substring '%v' but it was not found in the result", subs)
				}
			}

			if tt.input > 0 {
				emailPrefix := strings.Split(email, "@")[0]
				if len(emailPrefix) != tt.input {
					t.Errorf("Length of email prefix: %d, expecting %d", len(emailPrefix), tt.input)
				}
			}
			t.Log("Finished test:", tt.name)
		})
	}
}


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Valid Number", 5, 5},
		{"Zero Input", 0, 0},
		{"Negative Input", -3, 0},
		{"Long String", 1000, 1000},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Scenario:", tt.name)

			result := RandomString(tt.input)
			resultLength := len(result)

			if resultLength != tt.expected {
				t.Errorf("Failure in %s, expected length %v, but got %v", tt.name, tt.expected, resultLength)
			} else {
				t.Logf("Success in %s, expected length matches the actual length", tt.name)
			}
		})
	}
}

func main() {
	fmt.Println("Random String of length 5:", RandomString(5))
}

