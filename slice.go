package main

import "fmt"

func main() {
	//membuat array dengan jumlah elemen yang tidak ditentukan
	laptopBrands := [...]string{"Asus", "Acer", "Lenovo", "HP", "Dell", "Axioo", "Zyrex", "Infinix"}
	slice := laptopBrands[:7]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	//membuat slice dengan jumlah elemen dinamis
	drinkBrands := []string{"Coca Cola"}
	fmt.Println(drinkBrands)

	//Menambahkan elemen baru ke dalam slice drinkBrands
	drinkBrands = append(drinkBrands, "Fanta", "Sprite", "Pepsi")
	fmt.Println(drinkBrands)
	fmt.Println(drinkBrands[3])

	fmt.Println(len(drinkBrands)) //jumlah elemen di dalam slice
	fmt.Println(cap(drinkBrands)) //kapasitas slice

	mouseBrands := make([]string, 5, 10) //membuat slice dengan panjang 5 dan kapasitas 10
	fmt.Println(mouseBrands)

	mouseBrands[0] = "Logitech"
	mouseBrands[1] = "Razer"

	fmt.Println(mouseBrands)

	Brands := make([]string, len(mouseBrands), cap(mouseBrands)) //membuat slice baru dengan panjang dan kapasitas sama dengan mouseBrands
	fmt.Println(Brands)

	copy(Brands, mouseBrands) //menyalin isi slice mouseBrands ke dalam slice Brands
	fmt.Println(Brands)
}
