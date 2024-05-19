package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)


// Hello returns a greeting of  the named person
func Hello(name string) (string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}
	// Create a random message using a random format
	message:=fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string{
	// A slice of message formats.
	formats := []string{
		"Hi %v, welcome",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice format
	return formats[rand.Intn(len(formats))]
}
