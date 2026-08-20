package main

import "fmt"

type Address struct {
	City     string
	Province string
	Region   string
}

func changeAddress(address *Address) {
	address.Region = "Indonesia"
	fmt.Println(address)
}

func main() {
	address := Address{"Subang", "Jawa Barat", ""}
	changeAddress(&address)

	fmt.Println(address)
}
