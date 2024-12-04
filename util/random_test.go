package util_test

import (
	"strings"
	"testing"
	"github.com/wil-ckaew/gofinance-backend/util"
	"math/rand"
	"time"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	testCases := []struct {
		name           string
		input          int
		expectedLen    int
		expectedSuffix string
	}{
		{
			"RandomEmail creation with valid length",
			10,
			19,
			"@email.com",
		},
		{
			"RandomEmail creation with length 0",
			0,
			9,
			"@email.com",
		},
		{
			"RandomEmail creation with negative length",
			-1,
			9,
			"@email.com",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			result := RandomEmail(tt.input)

			if len(result) != tt.expectedLen {
				t.Errorf("Expected length %d, but got %d", tt.expectedLen, len(result))
			}
			if !strings.HasSuffix(result, tt.expectedSuffix) {
				t.Errorf("Expected suffix %s, but got %s", tt.expectedSuffix, result[len(result)-len(tt.expectedSuffix):])
			}
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name          string
		input         int
		expectedError bool
	}{
		{"Test RandomString to generate a string of given length", 5, false},
		{"Test RandomString with zero as input parameter", 0, false},
		{"Test RandomString with negative number as input parameter", -1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := util.RandomString(tc.input)

			switch {
			case tc.expectedError && res != "":
				t.Errorf("FAIL: Expected error but received string %s", res)
			case !tc.expectedError && res == "":
				t.Errorf("FAIL: Expected string of non-zero length but got an empty string")
			case len(res) != tc.input:
				t.Errorf("FAIL: Expected string of length %d but got string of length %d", tc.input, len(res))
			default:
				t.Logf("PASS: Received string of length %d", len(res))
			}
		})
	}

	str1 := util.RandomString(10)
	str2 := util.RandomString(10)
	if str1 == str2 {
		t.Errorf("FAIL: Expected different strings but got the same %s and %s", str1, str2)
	} else {
		t.Logf("PASS: Received two different strings %s and %s", str1, str2)
	}
}

