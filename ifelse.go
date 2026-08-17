package main

import "fmt"

func main() {
	name := "insan"

	//else if statement
	if name == "agus" {
		fmt.Println("Good morning, Insan")
	} else if name == "tony" {
		fmt.Println("Good morning, Agus")
	} else {
		fmt.Println("may i know your name?")
	}

	//if statement with short statement
	if length := len(name); length < 5 {
		fmt.Println("your name is too short")
	} else {
		fmt.Println("your name is long enough")
	}
}
