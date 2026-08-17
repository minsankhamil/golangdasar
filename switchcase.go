package main

import "fmt"

func main() {
	hariIni := "Jumat"
	//switchcase
	switch hariIni {
	case "Jumat":
		fmt.Println("Jam kerja hari ini dimulai jam 09.00 s/d 15.00")
	case "Sabtu", "Minggu":
		fmt.Println("Libur, tidak ada kegiatan bekerja")
	default:
		fmt.Println("Jam kerja hari ini dimulai jam 09.00 s/d 16.30")
	}

	ID := "Anto"
	//switchcase dengan short statement
	switch length := len(ID); length > 8 {
	case true:
		fmt.Println("ID sudah benar")
	case false:
		fmt.Println("ID terlalu pendek")
	}
}
