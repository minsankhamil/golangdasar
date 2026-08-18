package main

import "fmt"

//funvtion with named return values
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Antonius"
	middleName = "Adrian"
	lastName = "Sukajaya"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0
	//_ adalah indeks yang di blank kan
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
