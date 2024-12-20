package util

import (
	"fmt"
	"os"
	"io"
	"strings"
	"testing"
	"regexp"
	"math/rand"
	"time"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	cases := []testCase{
		{6, `.+@email.com`, false},
		{0, `@email.com`, false},
		{-1, "", true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input:%v", tc.input), func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			res := RandomEmail(tc.input)

			w.Close()
			os.Stdout = old

			var buf strings.Builder
			io.Copy(&buf, r)

			match, err := regexp.MatchString(tc.expected, res)
			if err != nil {
				t.Fatal(err)
			}

			if !match && !tc.hasError {
				t.Errorf("Expected email to match regex (%v), got %v ", tc.expected, res)
			} else if tc.hasError && match {
				t.Errorf("Expected error, got valid email %v", res)
			} else {
				t.Logf("success: expected and got match with regex (%v) ", tc.expected)
			}

			if unexpectedOutput := buf.String(); unexpectedOutput != "" {
				t.Errorf("unexpected print to stdout: %v", unexpectedOutput)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		input    int
		expected int
	}{
		{

			input:    5,
			expected: 5,
		}, {

			input:    0,
			expected: 0,
		}, {

			input:    -1,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {

			result := RandomString(tc.input)

			if len(result) != tc.expected {
				t.Logf("For input %v, expected length %v but got %v", tc.input, tc.expected, len(result))
				t.Fail()
			} else {
				t.Logf("Success scenario for input: %v", tc.input)
			}
		})
	}
}

