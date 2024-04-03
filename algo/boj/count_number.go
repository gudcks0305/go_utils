package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 리스트 크기 입력받기
	scanner.Scan()
	v, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 리스트 값 입력받기
	scanner.Scan()
	listStr := scanner.Text()
	listValues := strings.Split(listStr, " ")
	list := make([]int, v)
	for i := 0; i < v; i++ {
		list[i], err = strconv.Atoi(listValues[i])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// 찾을 값 입력받기
	scanner.Scan()
	findIntVar, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 찾을 값 카운트
	count := 0
	for _, val := range list {
		if val == findIntVar {
			count++
		}
	}
	fmt.Println(count)
}
