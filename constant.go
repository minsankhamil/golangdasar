package main

import "fmt"

func main() {
	//Data pada konstanta tidak dapat diubah nilainya, jika diubah maka akan terjadi error
	const firstname string = "Muhammad"
	const lastname string = "Insan"

	fmt.Println(firstname, lastname)

	//membuat multiple konstanta sekaligus
	const (
		firstname2 = "Aji"
		lastname2  = "Ramadhan"
	)

	fmt.Println(firstname2, lastname2)
}
