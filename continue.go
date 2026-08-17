package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			//Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
