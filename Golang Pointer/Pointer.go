package main

import "fmt"

func main() {
	var x int
	x = 3
	fmt.Println("x = ", x, "address of x = ", &x)
	//define pointer
	var y *int
	y = &x //assign pointer y with address variable x
	fmt.Println("y = ", *y, "address of y = ", y)
	*y = 4
	fmt.Println("Value y after change = ", *y)
	fmt.Println("Value x After y change:", x)
}
