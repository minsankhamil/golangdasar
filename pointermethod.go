package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) Married() {
	person.Name = "Mr. " + person.Name
	fmt.Println("Selamat Menempuh Hidup Baru", person.Name)
}

func main() {
	agi := Person{"Agi"}
	agi.Married()
	fmt.Println("Selamat Menempuh Hidup Baru", agi.Name)
}
