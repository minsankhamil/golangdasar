package main

import "fmt"

func data() interface{} {
	return "OK"
}

func main() {
	result := data()
	resultString := result.(string)
	fmt.Println(resultString)

	//resultInt := result.(int) //panic
	//fmt.Println(resultInt)
}
