package main

import "fmt"

func main() {
	var NilaiAkhir = 85
	var Absensi = 16

	var lulusNilaiAkhir bool = NilaiAkhir >= 80
	var lulusAbsensi bool = Absensi >= 12

	var lulusMataKuliah bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulusMataKuliah)
}
