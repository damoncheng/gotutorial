package main

import (
	"fmt"
	"github.com/damoncheng/gotutorial/pkg/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	// message, err := greetings.Hello("Gladys")
	names := []string{"Gladys", "Samantha", "Darrin"}

	//Request greeting messages for the names
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("success")
	fmt.Println(messages)
}
