package util

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	scenarios := []struct {
		name          string
		number        int
		resultPattern string
		hasError      bool
	}{
		{
			name:          "RandomEmail with a valid number",
			number:        5,
			resultPattern: "*****@email.com",
			hasError:      false,
		},
		{
			name:          "RandomEmail with zero as a number",
			number:        0,
			resultPattern: "@email.com",
			hasError:      false,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			email := RandomEmail(tt.number)
			emailPrefix := strings.Split(email, "@")[0]
			if len(emailPrefix) != tt.number {
				t.Errorf("Generated string length '%d' and expected string length '%d' do not match", len(emailPrefix), tt.number)
			} else if tt.hasError && len(emailPrefix) == tt.number {
				t.Errorf("Should throw an error for invalid number '%d'", tt.number)
			}
			t.Logf("Scenario passed: %s", tt.name)
		})
	}
}

