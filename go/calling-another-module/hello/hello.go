package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	// get a greeting message and print it.
	message := greetings.Hello("Eros")
	fmt.Println(message)
}
