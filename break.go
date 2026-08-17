package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			//Break digunakan untuk menghentikan seluruh perulangan
			break
		}
		fmt.Println("Perulangan ke", i)
	}

}
