package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name
func TestHelloName(t *testing.T) {
	name := "Sanjib"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// TestHelloName calls greetings.Hello with emtpy string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
