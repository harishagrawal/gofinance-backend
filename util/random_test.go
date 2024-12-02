package util

import (
	"fmt"
	"strings"
	"testing"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "String Zero Length",
			input:  0,
			expect: 0,
		},
		{
			name:   "String Length One",
			input:  1,
			expect: 1,
		},
		{
			name:   "Normal String Length",
			input:  50,
			expect: 50,
		},
		{
			name:   "Large String Length",
			input:  10000,
			expect: 10000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := RandomString(tc.input)

			if len(output) != tc.expect {
				t.Errorf("%s: expected %v, but got %v", tc.name, tc.expect, len(output))
			} else {
				t.Logf("%s: expected %v, got %v", tc.name, tc.expect, len(output))
			}
		})
	}
}

