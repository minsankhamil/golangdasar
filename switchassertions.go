package main

import "fmt"

func data() interface{} {
	return 2
}

func main() {
	result := data()
	switch value := result.(type) {
	case string:
		fmt.Println("This value =", value, ", is string")
	case int:
		fmt.Println("This value =", value, ", is int")
	default:
		fmt.Println("This value =", value, ", is unknown")
	}
}
