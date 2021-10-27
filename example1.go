package main

import "fmt"

// declare function type
type Operation func(balance, cantidad int) int


// function which its return type is an Operation (which is also a function)
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
