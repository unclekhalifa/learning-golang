package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

/*
	In order to use local greetings module, we use the following command
	go mod edit -replace example.com/greetings=../greetings
*/
func main() {
	log.SetPrefix("greetings: ") // Prefix for the logger
	log.SetFlags(0)              // Set flag to disable printing time, source file and line number

	names := []string{"John", "Ayo", "Mack"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
