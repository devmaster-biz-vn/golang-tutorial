package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func swap(a, b int) (int, int) {
	return b, a
}

func total(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	a := 10
	b := 5
	c := 7
	d := 8
	e := 11
	m, n := swap(a, b)
	fmt.Println("Swap a and b: ", m, n)
	fmt.Println("Add between a and b: ", add(a, b))
	fmt.Println("Subtract between a and b: ", sub(a, b))
	fmt.Println("Multiply between a and b: ", mul(a, b))
	fmt.Println("Divide between a and b:", div(a, b))
	fmt.Println("Total all: ", total(a, b, c, d, e))
}
