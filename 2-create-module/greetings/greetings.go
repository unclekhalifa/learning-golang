package greetings

import (
	"errors"

	"fmt"
)

/*
	Returns a greeting for the named person

	In Go, a fxn whose name starts with a capital letter
	can be called by a fxn not in the same package.
	This is known as an exported name
*/
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
