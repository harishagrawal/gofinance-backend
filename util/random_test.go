package util

import (
	rand "math/rand"
	strings "strings"
	testing "testing"
	time "time"
	debug "runtime/debug"
)








/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd

FUNCTION_DEF=func RandomEmail(number int) string 

*/
func TestRandomEmail(t *testing.T) {

	getLocal := func(email string) string {
		i := strings.Index(email, "@")
		if i == -1 {
			return ""
		}
		return email[:i]
	}

	isValidLocalPart := func(s string) bool {
		for _, r := range s {
			if !(('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')) {
				return false
			}
		}
		return true
	}

	type scenario struct {
		name          string
		input         int
		expectedLocal int
		expectExact   string
		expectValid   bool
		expectUnique  bool
		runs          int
		comment       string
		seeded        bool
		seed          int64
	}

	tests := []scenario{
		{
			name:          "Positive length, normal (Scenario 1)",
			input:         10,
			expectedLocal: 10,
			expectValid:   true,
			comment:       "Should generate a 10-char email local-part with valid format.",
		},
		{
			name:        "Zero length (Scenario 2)",
			input:       0,
			expectExact: "@email.com",
			comment:     "Should return exactly '@email.com' if input is 0.",
			expectValid: true,
		},
		{
			name:        "Negative length (Scenario 3)",
			input:       -5,
			expectExact: "@email.com",
			comment:     "Should return '@email.com' for negative input; check defensive handling.",
			expectValid: true,
		},
		{
			name:          "Large length, stress (Scenario 4)",
			input:         1000,
			expectedLocal: 1000,
			expectValid:   true,
			comment:       "Should handle large local part length, verify size/format.",
		},
		{
			name:          "Multiple calls uniqueness (Scenario 5)",
			input:         8,
			expectedLocal: 8,
			expectUnique:  true,
			runs:          10,
			expectValid:   true,
			comment:       "Should generate unique emails when called repeatedly.",
		},
		{
			name:          "No illegal chars in local part (Scenario 6)",
			input:         12,
			expectedLocal: 12,
			expectValid:   true,
			comment:       "Should only contain valid characters (letters/digits).",
		},
		{
			name:          "Minimal length (Scenario 7)",
			input:         1,
			expectedLocal: 1,
			expectValid:   true,
			comment:       "Should handle minimal length with correct formatting.",
		},
		{
			name:        "Various lengths (Scenario 8)",
			comment:     "Should work for a range of lengths.",
			expectValid: true,
			runs:        4,
		},
		{
			name:          "Seeded randomness - uniqueness over time (Scenario 9)",
			input:         8,
			expectedLocal: 8,
			expectUnique:  true,
			expectValid:   true,
			seeded:        true,
			runs:          2,
			comment:       "Should generate different emails with different seeds.",
		},
		{
			name:          "No struct dependency (Scenario 10)",
			input:         6,
			expectValid:   true,
			expectedLocal: 6,
			comment:       "Should not require Claims struct; test independence.",
		},
	}

	var varietyInputs = []int{2, 5, 15, 25}

	for _, tst := range tests {

		t.Run(tst.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Log("Scenario:", tst.comment)

			switch {
			case tst.runs > 1 && tst.name == "Multiple calls uniqueness (Scenario 5)":

				emails := make(map[string]struct{})
				for i := 0; i < tst.runs; i++ {
					email := RandomEmail(tst.input)
					local := getLocal(email)
					if len(local) != tst.expectedLocal {
						t.Errorf("Failed: Expected local part length %d, got %d in '%s'", tst.expectedLocal, len(local), email)
					}
					if !strings.HasSuffix(email, "@email.com") {
						t.Errorf("Failed: Expected suffix '@email.com', got '%s'", email)
					}
					if !isValidLocalPart(local) {
						t.Errorf("Failed: Local part '%s' contains illegal characters.", local)
					}
					emails[email] = struct{}{}
				}
				if len(emails) < tst.runs {
					t.Errorf("Failed: RandomEmail output not unique for %d runs.", tst.runs)
				} else {
					t.Log("Success: All generated emails are unique.")
				}

			case tst.name == "Various lengths (Scenario 8)":
				for _, inputLen := range varietyInputs {
					email := RandomEmail(inputLen)
					local := getLocal(email)
					if len(local) != inputLen {
						t.Errorf("Failed: Expected local part len %d, got %d in '%s'", inputLen, len(local), email)
					}
					if !strings.HasSuffix(email, "@email.com") {
						t.Errorf("Failed: Email for input %d should end with '@email.com', got '%s'", inputLen, email)
					}
					if !isValidLocalPart(local) {
						t.Errorf("Failed: Local part for input %d has illegal chars: '%s'", inputLen, local)
					}
					t.Logf("Success: Input %d generated valid email '%s'", inputLen, email)
				}

			case tst.seeded && tst.name == "Seeded randomness - uniqueness over time (Scenario 9)":
				emails := []string{}
				seeds := []int64{time.Now().UnixNano(), time.Now().UnixNano() + 5000}
				for i := 0; i < tst.runs; i++ {
					rand.Seed(seeds[i])
					email := RandomEmail(tst.input)
					emails = append(emails, email)
					local := getLocal(email)
					if len(local) != tst.expectedLocal {
						t.Errorf("Failed: Seeded run %d: Expected local part length %d, got %d", i, tst.expectedLocal, len(local))
					}
					if !strings.HasSuffix(email, "@email.com") {
						t.Errorf("Failed: Seeded run %d: Expected suffix '@email.com', got '%s'", i, email)
					}
					if !isValidLocalPart(local) {
						t.Errorf("Failed: Seeded run %d: Local part contains illegal characters in '%s'", i, email)
					}
				}
				if emails[0] == emails[1] {
					t.Errorf("Failed: Seeded randomness produced duplicate emails: '%s'", emails[0])
				} else {
					t.Log("Success: Different seeds produced unique emails.")
				}

				rand.Seed(time.Now().UnixNano())

			default:

				email := RandomEmail(tst.input)
				local := getLocal(email)

				if tst.expectExact != "" {
					if email != tst.expectExact {
						t.Errorf("Failed: Expected '%s', got '%s'", tst.expectExact, email)
					} else {
						t.Logf("Success: Exact match '%s'", email)
					}
				}
				if tst.expectedLocal >= 0 && tst.expectExact == "" {
					if len(local) != tst.expectedLocal {
						t.Errorf("Failed: Expected local part length %d, got %d ('%s')", tst.expectedLocal, len(local), local)
					}
				}
				if !strings.HasSuffix(email, "@email.com") {
					t.Errorf("Failed: Email '%s' does not end with '@email.com'", email)
				}
				if tst.expectValid && !isValidLocalPart(local) {
					t.Errorf("Failed: Local part '%s' contains illegal characters.", local)
				}
				t.Log("Success: Email format/length valid.")

				if tst.name == "No struct dependency (Scenario 10)" {

					t.Log("Success: No Claims struct interaction required for RandomEmail.")
				}
			}
		})
	}
}

