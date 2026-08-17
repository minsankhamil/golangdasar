package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yusuf",
		"age":     "20",
		"address": "Jakarta",
	}
	fmt.Println(person)
	fmt.Println(person["name"])

	fmt.Println(len(person)) //jumlah elemen di dalam map
	fmt.Println(person["age"])

	delete(person, "age") //menghapus elemen dengan key "age"
	fmt.Println(person)

	person["title"] = "junior programmer" //menambahkan elemen baru dengan key "title"
	person["address"] = "Bandung"         //mengubah nilai elemen dengan key "address"
	fmt.Println(person)
}
