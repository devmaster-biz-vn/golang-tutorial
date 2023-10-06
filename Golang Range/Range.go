package main

import "fmt"

func main() {
	//Use range with array
	arr := [3]string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	//Use range with map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	//Use range with slice
	slice := []int{4, 5, 6, 8, 0, 1}
	for index, value := range slice {
		fmt.Println("Index: ", index, "Value: ", value)
	}
}
