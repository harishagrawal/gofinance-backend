package null

import (
	"testing"
	"time"
	"math/rand"
)

/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name       string
		input      int
		wantLength int
		wantEmail  string
	}{
		{
			name:       "Normal Case - Generating a random 7 characters long email.",
			input:      7,
			wantLength: 17,
		},
		{
			name:       "Edge Case - Generating a random 0 characters long email.",
			input:      0,
			wantLength: 10,
			wantEmail:  "@email.com",
		},
		{
			name:       "Error Case - Negative Input.",
			input:      -4,
			wantLength: 10,
			wantEmail:  "@email.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			gotEmail := RandomEmail(tc.input)

			if len(gotEmail) != tc.wantLength {
				t.Errorf("got the length of email as %d; want %d", len(gotEmail), tc.wantLength)
			}

			if tc.wantEmail != "" && gotEmail != tc.wantEmail {
				t.Errorf("got email as %s; want %s", gotEmail, tc.wantEmail)
			}
		})
	}
}

