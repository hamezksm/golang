package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main (){
	// Set properties of predefined Logger, including the entry
	// prefix of a flag to disable printing the time, source and 
	// line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"James", "John", "Janet"}

	// Request greeting messages for names.
	messages, err := greetings.Hellos(names)
	// If error was returned, print to console and exit the program
	if err != nil{
		log.Fatal(err)
	}
	// If no error was returned, print returned message to console
	fmt.Println(messages)
}