package main

import "fmt"

// CalculateSum обчислює суму двох цілих чисел відповідно до стандарту gofmt.
func CalculateSum(a int, b int) int {
	return a + b
}

func main() {
	result := CalculateSum(10, 25)
	fmt.Println("Result:", result)
}
