package main

import "fmt"

func main() {
	//define a array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] //slice  array from index 1 to index 3
	fmt.Println(slice)

	slice1 := arr[1:] //slice array from index 1 to end
	fmt.Println(slice1)

	slice2 := arr[:3] //slice array from begin to index 2
	fmt.Println(slice2)
}
