package main

import "fmt"

func main() {
	//membuat variable, diikuti nama variable, diikuti tipe data
	//data pada variable dapat diubah-ubah
	var name string

	name = "Muhammad Insan"
	fmt.Println(name)

	name = "Insan Khamil"
	fmt.Println(name)

	//membuat variable, diikuti nama variable, tanpa menyebutkan tipe data
	var name2 = "Ananda Yasmin"
	fmt.Println(name2)

	name2 = "Fani Munaroh"
	fmt.Println(name2)

	// lebih ringkas dengan kata kunci :=
	name3 := "Afrianto Sangir"
	fmt.Println(name3)

	// membuat multiple variable sekaligus
	var (
		firstName = "Muhammad"
		lastName  = "Insan"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName, lastName)
}
