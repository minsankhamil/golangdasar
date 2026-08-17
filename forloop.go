package main

import "fmt"

func main() {
	counter := 1

	for counter <= 50 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for counter := 2; counter <= 20; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	brands := []string{"Apple", "Samsung", "Xiaomi", "Oppo"}
	for x, name := range brands {
		fmt.Println("index", x, "=", name)
	}
}
