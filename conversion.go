package main

import "fmt"

func main() {
	//konversi tipe data 1
	var nilai1 int32 = 33100
	var nilai2 int64 = int64(nilai1)

	var nilai3 int16 = int16(nilai1)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)

	var name = "Muhammad Insan"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
