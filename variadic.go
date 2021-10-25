package main

import "fmt"

//Variadic function
func promedio(marks ...int) int {
	var promedio, suma int
	for _, mark := range marks {
		suma += mark
	}
	promedio = suma / len(marks)
	return promedio
}


func main() {
	//result = promedio(10)
	//result = promedio(10, 10)
	//result = promedio(10, 10, 10, 10)
	nums := []int{1, 2, 3, 4}
	result = promedio(nums...)
	fmt.Println("variadic result:", result)
}
