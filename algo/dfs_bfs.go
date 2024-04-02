package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	arr     [][]int
	visited []bool
	number  int
)

// visited를 2번 사용하니까 리셋시켜주기.
func visitedInitialized() {
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
}

func dfs(v int) {
	visited[v] = true
	fmt.Print(v, " ")

	for i := 1; i <= number; i++ {
		if arr[v][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}

func bfs(v int) {
	visited[v] = true
	queue := []int{v}

	for len(queue) != 0 {
		front := queue[0]
		fmt.Print(front, " ")
		queue = queue[1:]

		for i := 1; i <= number; i++ {
			if arr[front][i] == 1 && !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer reader.Reset(os.Stdin)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}(writer)

	var m, startV int
	_, err := fmt.Fscanln(reader, &number, &m, &startV)
	if err != nil {
		panic(err)
	}

	arr = make([][]int, number+1)
	for i := range arr {
		arr[i] = make([]int, number+1)
	}

	visited = make([]bool, number+1)

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		arr[x][y] = 1
		arr[y][x] = 1
	}

	dfs(startV)
	visitedInitialized()
	fmt.Println()

	bfs(startV)

}
