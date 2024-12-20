package util

import (
	"testing"
	"math/rand"
	"time"
	"strings"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name       string
		input      int
		wantLength int
		wantSuffix string
	}{
		{
			name:       "Valid Email Generation",
			input:      10,
			wantLength: 21,
			wantSuffix: "@email.com",
		},
		{
			name:       "Zero Length Email",
			input:      0,
			wantLength: 11,
			wantSuffix: "@email.com",
		},
		{
			name:       "Negative Length Email",
			input:      -5,
			wantLength: 11,
			wantSuffix: "@email.com",
		},
		{
			name:       "Large Length Email",
			input:      10000,
			wantSuffix: "@email.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomEmail(tt.input)

			if tt.wantLength != 0 && len(got) != tt.wantLength {
				t.Fatalf("RandomEmail() length = %v, want %v", len(got), tt.wantLength)
			}

			if !strings.HasSuffix(got, tt.wantSuffix) {
				t.Fatalf("RandomEmail() suffix = %v, want %v", got[len(got)-10:], tt.wantSuffix)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	scenarios := []struct {
		desc     string
		input    int
		expected int
	}{
		{
			desc:     "Testing RandomString with valid number",
			input:    5,
			expected: 5,
		},
		{
			desc:     "Testing RandomString with zero",
			input:    0,
			expected: 0,
		},
		{
			desc:     "Testing RandomString with negative number",
			input:    -5,
			expected: 0,
		},
		{
			desc:     "Testing RandomString with a large number",
			input:    1000000,
			expected: 1000000,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			res := RandomString(s.input)
			if len(res) != s.expected {
				t.Errorf("failure in scenario: %s: expected %d, got %d", s.desc, s.expected, len(res))
			}
		})
	}
}

