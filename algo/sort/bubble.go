package main

import "fmt"
func main() {
	arr := [...]int{5,2,1,8,9}
	fmt.Println(bubbleSort(arr[:]))
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}