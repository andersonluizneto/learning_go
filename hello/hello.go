package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Looger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, ande line number.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greetings message.
	message, err := greetings.Hello(strings.Join(os.Args[1:], " "))

	// If an error was returned, print it to the cosole and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print it to the returned message
	// to the console.
	fmt.Println(message)
}
