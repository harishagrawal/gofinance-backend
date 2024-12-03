package util_test

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
func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{
			name:    "Generate valid random email for the number greater than zero",
			length:  10,
			wantErr: false,
		},
		{
			name:    "Generate valid random email for the number equals to zero",
			length:  0,
			wantErr: false,
		},
		{
			name:    "Generate valid random email for a very large number",
			length:  100000,
			wantErr: false,
		},
		{
			name:    "Generate valid random email for negative numbers",
			length:  -1,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := RandomEmail(tc.length)
			parts := strings.Split(email, "@")

			if len(parts) != 2 || parts[1] != "email.com" {
				t.Fatalf("Unexpected email format: %s", email)
			}

			if tc.length >= 0 && len(parts[0]) != tc.length {
				t.Fatalf("Expected length %d, got %d", tc.length, len(parts[0]))
			}

			if tc.length < 0 && !tc.wantErr {
				t.Fatalf("Expected error for negative input, but function returned valid result")
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func RandomString(number int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomString(t *testing.T) {

	var alphabet = "abcdefghijklmnopqrstuvwxyz"

	rand.Seed(time.Now().UnixNano())

	cases := []struct {
		name string
		num  int
	}{
		{
			name: "Random String Generation Test",
			num:  5,
		},
		{
			name: "Zero Length String Test",
			num:  0,
		},
		{
			name: "Negative Length String Test",
			num:  -5,
		},
		{
			name: "Random String Qualitative Test",
			num:  10,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			res := RandomString(tt.num)

			if tt.num <= 0 {
				if res != "" {
					t.Errorf("expected empty string for non-positive input, got %q", res)
				}
			} else {
				if len(res) != tt.num {
					t.Errorf("expected string of length %d, got string of length %d", tt.num, len(res))
				}
				for _, c := range res {
					if !strings.ContainsRune(alphabet, c) {
						t.Errorf("generated string contains unexpected character %q", c)
					}
				}
			}
		})
	}
}

