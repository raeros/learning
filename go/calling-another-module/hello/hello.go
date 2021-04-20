package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main(){

	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	// get a greeting message and print it.
	message, err := greetings.Hello("")

	// if an Error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no Error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
