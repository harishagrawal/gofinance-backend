package util

import (
	"strings"
	"testing"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	type test struct {
		input  int
		output string
		err    error
	}

	tests := []test{
		{5, "@email.com", nil},
		{0, "@email.com", nil},
		{1024, "@email.com", nil},
		{-5, "@email.com", nil},
	}

	for _, test := range tests {
		t.Logf("Executing test with input: %v", test.input)

		result := RandomEmail(test.input)

		if !strings.Contains(result, "@") {
			t.Fatalf("Expected email to contain '@' but it wasn't found")
		}

		lenEmail := len(result) - 10

		if lenEmail != test.input {
			t.Fatalf("Generated email with wrong length. Expected length: %v, Got: %v", test.input, lenEmail)
		}

		if lenEmail < 0 {
			t.Fatalf("Generated email length was less than zero for input: %v", test.input)
		}

		t.Logf("Successful test with input: %v", test.input)
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	cases := []struct {
		name             string
		length           int
		expectedLength   int
		expectInAlphabet bool
	}{
		{
			name:             "Random String Generation Test",
			length:           10,
			expectedLength:   10,
			expectInAlphabet: true,
		},
		{
			name:             "Zero Length String Test",
			length:           0,
			expectedLength:   0,
			expectInAlphabet: false,
		},
		{
			name:             "Negative Length String Test",
			length:           -5,
			expectedLength:   0,
			expectInAlphabet: false,
		},
		{
			name:             "Random String Qualitative Test",
			length:           20,
			expectedLength:   20,
			expectInAlphabet: true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			rs := RandomString(test.length)

			if len(rs) != test.expectedLength {
				t.Errorf("expected length %d, got %d, test case: %s", test.expectedLength, len(rs), test.name)
			}

			if test.expectInAlphabet {
				for i := range rs {
					if !strings.Contains(alphabet, string(rs[i])) {
						t.Errorf("unexpected character at %d, test case: %s", i, test.name)
					}
				}
			}
		})
	}
}

