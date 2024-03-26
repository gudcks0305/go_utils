package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 2)
	fmt.Println("myslice")
	fmt.Println(mySlice)

	destSlice := make([]int, len(mySlice))
	copy(destSlice, mySlice)
	fmt.Println("dest")
	fmt.Println(destSlice)

	sliceList := destSlice[1:2]
	fmt.Println("slice")
	fmt.Println(sliceList)
}
