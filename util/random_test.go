package util

import (
	"strings"
	"testing"
	"fmt"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	testCases := []struct {
		desc         string
		input        int
		expectSuffix string
		expectLength int
	}{
		{
			desc:         "Normal operation with valid input",
			input:        5,
			expectSuffix: "@email.com",
			expectLength: 15,
		},
		{
			desc:         "Edge case with zero as input",
			input:        0,
			expectSuffix: "@email.com",
			expectLength: 10,
		},
		{
			desc:         "Edge case with negative integer input",
			input:        -5,
			expectSuffix: "@email.com",
			expectLength: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			email := RandomEmail(tc.input)
			if !strings.HasSuffix(email, tc.expectSuffix) {
				t.Errorf("expected suffix %s, got %s", tc.expectSuffix, email)
			}

			if len(email) != tc.expectLength {
				t.Errorf("expected length %d, got %d", tc.expectLength, len(email))
			}

			t.Logf("Test case %s passed.", tc.desc)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().Unix())

	var testCases = []struct {
		length      int
		expectedLen int
		name        string
	}{
		{5, 5, "Positive Non Zero Integers as Input"},
		{0, 0, "Zero as Input"},
		{-5, 0, "Negative Integers as Input"},
		{10000, 10000, "Large Integers as Input"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RandomString(tc.length)

			if len(result) != tc.expectedLen {
				t.Errorf("RandomString(%d) failed with length %d, expected %d", tc.length, len(result), tc.expectedLen)
			} else {
				t.Logf("RandomString(%d) returned a string of expected length %d", tc.length, tc.expectedLen)
			}
		})
	}
}

