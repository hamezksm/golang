// Also, test functions take a pointer to the testing package's testing.T type as a parameter.
// You use this parameter's methods for reporting and logging from your test.

package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "James"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("James")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("James") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty String
// Checking for error

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err ==nil{
		t.Fatalf(`Hello("") = %q ,%v, want "", error`, msg, err)
	}
}
