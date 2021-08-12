package greetings

import (
	"errors"

	"fmt"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// Map to associate names with messages
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
	randomFormat returns one of a set of greetings messages. The returned
	message is selected at random
*/
func randomFormat() string {
	// Size can grow dynamically
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
