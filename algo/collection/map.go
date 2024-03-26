package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["tt"] = 1
	myMap["ti"] = 2

	fmt.Println(myMap)

	// get key
	for key := range myMap {
		fmt.Println(key)
	}

	// get value

	for _, value := range myMap {
		fmt.Println(value)
	}
}
