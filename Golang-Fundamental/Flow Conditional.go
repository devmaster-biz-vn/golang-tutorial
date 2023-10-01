package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Print("Input Student Score: ")
	fmt.Scanln(&score)
	fmt.Println("If-Else conditional")
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
	fmt.Println("Switch-Case")
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
		fmt.Println("Try More")
	case score >= 60:
		fmt.Println("D")
		fmt.Println("Not Good")
	default:
		fmt.Println("E")
		fmt.Println("Very Bad")
	}
	fmt.Println("Loop")
	fmt.Println("Loop For")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("Loop While")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("Do While")
	x := 0
	for {
		fmt.Println(x)
		x++
		if x > 5 {
			break
		}
	}
}
