package main

import "fmt"

func main() {
	name1 := "Insan"
	name2 := "Khamil"
	name3 := "Insan"

	var result bool = name1 == name2
	var result2 bool = name1 == name3
	var result3 bool = name1 != name2

	fmt.Println("Apakah name1 sama dengan name2?", result)
	fmt.Println("Apakah name1 sama dengan name3?", result2)
	fmt.Println("Apakah name1 tidak sama dengan name2?", result3)
}
