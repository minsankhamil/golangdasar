package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Cahyo"
	names[1] = "Budi"
	names[2] = "Andi"

	fmt.Println(names)
	fmt.Println(names[1])

	var products = [3]string{"Laptop", "Mouse", "Keyboard"}
	fmt.Println(products)

	values := [3]int{5, 3, 9}
	fmt.Println(values)

	//jumlah elemen di dalam array
	fmt.Println(len(values))

	//mengubah nilai elemen array
	names[1] = "Anto"
	fmt.Println(names[1])
}
