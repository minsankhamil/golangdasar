package main

import "fmt"

type Address struct {
	City     string
	Province string
	Region   string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"
	address3.City = "Semarang"

	address2 = &Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}
	*address3 = Address{"Semarang", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
