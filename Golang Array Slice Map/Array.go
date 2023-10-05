package main

import "fmt"

func main() {
	//define arr with size
	var number [2]int
	number[0] = 1
	number[1] = 2

	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

	number1 := [5]int{4, 5, 6, 1, 9}
	for i := 0; i < len(number1); i++ {
		fmt.Println(number1[i])
	}

	number2 := [...]int{2, 5, 6, 82, 9, 45}
	fmt.Println(number2)

	fmt.Println("Loop Array with range:")
	for index, value := range number1 {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	//MultiDimensional Arrays
	var matrix [3][3]int
	matrix[0][0] = 0
	matrix[0][1] = 1
	matrix[0][2] = 2

	matrix[1][0] = 3
	matrix[1][1] = 4
	matrix[1][2] = 5

	matrix[2][0] = 6
	matrix[2][1] = 7
	matrix[2][2] = 8

	//Loop MultiDimensional Arrays
	for _, row := range matrix {
		fmt.Println()
		for _, value := range row {
			fmt.Print(value, " ")
		}
	}

}
