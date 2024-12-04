package util

import (
	"fmt"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	t.Run("random email with regular length", func(t *testing.T) {
		const length = 10
		t.Log("Random email generation with regular length scenario")

		email := RandomEmail(length)

		local, domain := splitEmail(email)

		assert.Equal(t, length, len(local), fmt.Sprintf("expected length %d, got %d", length, len(local)))
		assert.Equal(t, "@email.com", domain, "incorrect domain part, expecting @email.com")
	})

	t.Run("random email with zero length", func(t *testing.T) {
		const length = 0
		t.Log("Random email generation with zero length scenario")

		email := RandomEmail(length)

		local, domain := splitEmail(email)

		assert.Empty(t, local, fmt.Sprintf("expected length %d, got %d", length, len(local)))
		assert.Equal(t, "@email.com", domain, "incorrect domain part, expecting @email.com")
	})

	t.Run("random email with negative length", func(t *testing.T) {
		const length = -1
		t.Log("Random email generation with negative length scenario")

		email := RandomEmail(length)

		local, domain := splitEmail(email)

		assert.Empty(t, local, fmt.Sprintf("expected length to be 0 for negative input, got %d", len(local)))
		assert.Equal(t, "@email.com", domain, "incorrect domain part, expecting @email.com")
	})

	t.Run("random email with a very large length", func(t *testing.T) {
		const length = 1000
		t.Log("Random email generation with a very large length scenario")

		email := RandomEmail(length)

		local, domain := splitEmail(email)

		assert.Equal(t, length, len(local), fmt.Sprintf("expected length %d, got %d", length, len(local)))
		assert.Equal(t, "@email.com", domain, "incorrect domain part, expecting @email.com")
	})
}

func splitEmail(email string) (local, domain string) {
	at := strings.Index(email, "@")
	local = email[:at]
	domain = email[at:]
	return
}

/*
ROOST_METHOD_HASH=RandomString_d7e3599ac4
ROOST_METHOD_SIG_HASH=RandomString_c6fe4ad19a


 */
func TestRandomString(t *testing.T) {

	testCases := []struct {
		input int
		empty bool
	}{
		{20, false},
		{-5, true},
		{0, true},
		{10, false},
		{0, true},
		{10, false},
	}

	for i, tc := range testCases {

		result := RandomString(tc.input)

		require.Equal(t, tc.input <= 0, tc.empty, fmt.Sprintf("Case %d failed: Expected string length to be non-negative but got negative length", i+1))

		if tc.input > 0 && tc.empty != true {
			secondResult := RandomString(tc.input)
			require.NotEqual(t, result, secondResult, fmt.Sprintf("Case %d failed: random strings must not be equal", i+1))
		}

		t.Logf("Test case number %d executed successfully", i+1)
	}
}

