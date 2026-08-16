package main

import "fmt"

func main() {
	//Operasi Matematika
	a := 10
	b := 30
	c := a + b
	fmt.Println("Jika a = 10 dan b = 30, maka hasil dari a + b adalah", c)

	d := a * b
	fmt.Println("Jika a = 10 dan b = 30, maka hasil dari a * b adalah", d)

	e := b / a
	fmt.Println("Jika a = 10 dan b = 30, maka hasil dari b / a adalah", e)

	f := a - b
	fmt.Println("Jika a = 10 dan b = 30, maka hasil dari a - b adalah", f)

	//Augmented Assignment
	g := 10
	g += 15

	fmt.Println(g)

	//Unary Operator
	h := 50
	h++
	fmt.Println(h)

	h--
	fmt.Println(h)
}
