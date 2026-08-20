package main

import "fmt"

type Address struct {
	City     string
	Province string
	Region   string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Region = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
