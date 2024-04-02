package main

import "math"

func main() {
	list := []int{1, 2, 3, 4, 5}
	println(findMax(&list))
}

func findMax(list *[]int) int {
	maxVal := math.MinInt32
	for _, v := range *list {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
