package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Andri", "Marwanto")

	result := getHello("Yasmin")
	fmt.Println(result)

	firstName1, lastName2 := getFullName()
	fmt.Println(firstName1, lastName2)

	//skip return value with _
	firstName2, _ := getFullName()
	fmt.Println(firstName2)
}

func sayHello() {
	fmt.Println("Hello World")
}

//function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

//function with return value
func getHello(name string) string {
	return "hello " + name
}

//function with return multiple value
func getFullName() (string, string) {
	return "Asep", "Sukatani"
}
