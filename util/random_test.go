package util

import (
	"strings"
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/your_github_username/your_project_directory/util"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var testCases = []struct {
		name        string
		charLength  int
		expectedErr error
	}{
		{"Valid number of characters", 10, nil},
		{"Zero characters", 0, nil},
		{"Negative number", -4, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			email := util.RandomEmail(tc.charLength)
			prefixLength := len(strings.Split(email, "@")[0])

			assert.Contains(t, email, "@")
			assert.Contains(t, email, ".com")
			assert.NotContains(t, email, " ")
			assert.Equal(t, prefixLength, tc.charLength)
		})
	}
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	tests := []struct {
		name             string
		input            int
		isEmptyExpected  bool
		isLengthExpected bool
		isRandomExpected bool
	}{
		{
			name:             "Successful Generation of RandomString with Length 10",
			input:            10,
			isEmptyExpected:  false,
			isLengthExpected: true,
			isRandomExpected: true,
		},
		{
			name:            "Testing RandomString with Length 0",
			input:           0,
			isEmptyExpected: true,
		},
		{
			name:            "Testing RandomString with a Negative Length",
			input:           -5,
			isEmptyExpected: true,
		},
		{
			name:             "Randomness of RandomString",
			input:            10,
			isRandomExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := RandomString(tt.input)
			if tt.isEmptyExpected && res != "" {
				t.Errorf("Expected an empty string, but got: %v", res)
			} else if tt.isLengthExpected && len(res) != tt.input {
				t.Errorf("Mismatched lengths, length expected: %v, length returned: %v", tt.input, len(res))
			} else if tt.isRandomExpected && tt.input > 0 {
				res1 := RandomString(tt.input)
				if res == res1 {
					t.Errorf("Expected different strings, but got: %v and %v", res, res1)
				}
			}
		})
	}
}

