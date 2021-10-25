package main

import "fmt"

type Operation func(balance, cantidad int) int

func crearOperation(tipo string) Operation {
	if tipo == "suma" {
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo ==  "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}
}

func main() {
	suma := crearOperation("suma")
	result := suma(10, 15)
	fmt.Println("result:", result)
}
