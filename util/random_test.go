package util

import (
	"strings"
	"testing"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name        string
		input       int
		expectedLen int
		errExpected bool
	}{
		{
			name:        "Standard Random Email Generation",
			input:       5,
			expectedLen: 5 + len("@email.com"),
			errExpected: false,
		},
		{
			name:        "Edge Case with Zero as Input",
			input:       0,
			expectedLen: len("@email.com"),
			errExpected: false,
		},
		{
			name:        "Edge Case with Negative Input",
			input:       -1,
			expectedLen: len("@email.com"),
			errExpected: false,
		},
		{
			name:        "Large Number Input",
			input:       10000,
			expectedLen: 10000 + len("@email.com"),
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomEmail(tc.input)
			if len(result) != tc.expectedLen {
				t.Errorf("unexpected length for %q, expected %q,size, got: %q", tc.name, tc.expectedLen, len(result))
			}
			if !strings.Contains(result, "@email.com") {
				t.Errorf("email format for %q not as expected, expected '@email.com', got: %q", tc.name, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func func TestRandomString(t *testing.T) {
	testCases := []struct {
		name    string
		args    int
		wantErr bool
	}{
		{
			name:    "Normal operation with valid positive integer",
			args:    7,
			wantErr: false,
		},
		{
			name:    "Boundary case with zero input",
			args:    0,
			wantErr: false,
		},
		{
			name:    "Negative integer input",
			args:    -8,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			want := ""

			got := RandomString(tc.args)

			switch {
			case tc.wantErr:
				if got != "" {
					t.Errorf("RandomString() = %v, want empty string", got)
				}
			default:
				if len(got) != tc.args {
					t.Errorf("RandomString() = %v, want %v", len(got), tc.args)
				}

				if got == want {
					t.Errorf("RandomString() = %v, want a string different from the previous one %v", got, want)
				}

				want = got
			}
		})
	}
}

