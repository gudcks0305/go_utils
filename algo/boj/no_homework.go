package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	list := make([]bool, 31)
	for i := 0; i < 28; i++ {
		scanner.Scan()
		index, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		list[index] = true
	}
	for i := 1; i < len(list); i++ {
		if list[i] == false {
			fmt.Println(i)
		}
	}

}
