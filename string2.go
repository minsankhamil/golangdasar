package main

import "fmt"

func main() {
	//menghitung jumlah karakter di String
	fmt.Println(len("Muhammad"))
	//Mengambil karakter pada posisi yang ditentukan dalam bentuk byte(uint8)
	fmt.Println("Muhammad Insan"[0])
	//Mengambil karakter pada posisi yang ditentukan dalam bentuk byte(uint8)
	//lalu dikonversi menjadi string
	fmt.Println(string("Muhammad Insan Khamil"[0]))
}
