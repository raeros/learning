package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// if no name was given, return an erros with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
