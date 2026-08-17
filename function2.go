package main

import "fmt"

//funvtion with named return values
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Antonius"
	middleName = "Adrian"
	lastName = "Sukajaya"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
