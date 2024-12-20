package util

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"strconv"
)


/*
ROOST_METHOD_HASH=RandomEmail_1905439733
ROOST_METHOD_SIG_HASH=RandomEmail_7a04f189fd


 */
func TestRandomEmail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	cases := []struct {
		name  string
		input int
	}{
		{"Positive Number", 5},
		{"Negative Number", -3},
		{"Zero", 0},
		{"Large Number", 9000},
		{"Randomness", 7},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Executing scenario: ", tt.name)
			email := RandomEmail(tt.input)

			t.Log("Generated Email: ", email)
			switch tt.name {
			case "Positive Number":
				if len(email) != tt.input+10 {
					t.Errorf("Expected length of email is %v, but got %v", tt.input+10, len(email))
				}
			case "Negative Number":
				if email != "@email.com" {
					t.Errorf("Expected @email.com for negative input but got %v", email)
				}
			case "Zero":
				if email != "@email.com" {
					t.Errorf("Expected @email.com for zero input but got %v", email)
				}
			case "Large Number":
				if len(email) != tt.input+10 {
					t.Errorf("Expected length of email is %v, but got %v", tt.input+10, len(email))
				}
			case "Randomness":
				secondEmail := RandomEmail(tt.input)
				t.Log("Generated Second Email: ", secondEmail)
				if email == secondEmail {
					t.Errorf("Expected different emails but got same email twice")
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

	var tests = []struct {
		name   string
		length int
	}{
		{name: "Scenario 1: Successful Random String Generation", length: 5},
		{name: "Scenario 2: Zero Length String", length: 0},
		{name: "Scenario 3: Length Argument Is Negative", length: -5},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			res := RandomString(tt.length)

			if tt.length < 0 {
				if len(res) != 0 {
					t.Errorf("Expected length of Random String to be 0, but got %d", len(res))
				}
				t.Log(tt.name + " Passed!")
			} else if len(res) != tt.length {
				t.Errorf("Expected length of Random String to be " + strconv.Itoa(tt.length) + ", but got " + strconv.Itoa(len(res)))
			} else {
				t.Log(tt.name + " Passed!")
			}

		})
	}
}

