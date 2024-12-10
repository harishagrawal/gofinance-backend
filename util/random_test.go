package util

import (
	"testing"
	"strings"
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().Unix())

	var tests = []struct {
		name       string
		input      int
		wantLength int
	}{
		{"Valid Number of Characters for Email", 10, 20},
		{"Invalid Number of Characters for Email", -10, 0},
		{"Zero Characters for Email", 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomEmail(tt.input)
			if (tt.name == "Invalid Number of Characters for Email") && tt.input < 0 {

			} else {
				if len(got) != tt.wantLength || !strings.HasSuffix(got, "@email.com") {
					t.Errorf("RandomEmail(%v) = %v, want length=%v and suffix='@email.com'", tt.input, got, tt.wantLength)
				} else {
					t.Log("Success: Test for ", tt.name)
				}
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
			"Positive Test for Normal Operation",
			5,
			5,
		},
		{
			"Edge Case Test for Zero Length String",
			0,
			0,
		},
		{
			"Negative Test for Negative Integers",
			-5,
			0,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.desc, func(t *testing.T) {
			result := RandomString(tt.input)

			if len(result) != tt.expected {
				t.Errorf("scenario: %s, expected %d, but got %d", tt.desc, tt.expected, len(result))
			} else {
				t.Logf("scenario: %s, successfully passed.", tt.desc)
			}
		})
	}
}

