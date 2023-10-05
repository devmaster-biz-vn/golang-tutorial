package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	//delete key in map
	delete(m, "a")
	//check key in mapping exist

	a, exists := m["a"]
	if exists {
		fmt.Println("Mapping with key a:", a)
	} else {
		fmt.Println("Not found key in mapping")
	}

}
