package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is ", customer.name)
}

func main() {
	var eko Customer
	eko.name = "Eko Agusyana"
	eko.address = "Jakarta, Indonesia"
	eko.age = 29

	fmt.Println(eko)

	joko := Customer{
		name:    "Joko",
		address: "Indonesia",
		age:     32,
	}
	fmt.Println(joko)

	//struct method
	rully := Customer{name: "Rully"}
	rully.sayHello()
}
