package main

import "fmt"
import greetings "github.com/damoncheng/gotugorial"

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
