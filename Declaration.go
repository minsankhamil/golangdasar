package main

import "fmt"

func main() {
	type NoKTP string
	//Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	//Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

	var ktpEko NoKTP = "1234567890"
	fmt.Println(ktpEko)
	fmt.Println(NoKTP("1234567890"))
}
