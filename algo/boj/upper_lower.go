package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	charList := []rune(str)
	for i, r := range charList {
		if unicode.IsLower(r) {
			charList[i] = unicode.ToUpper(r)
		} else {
			charList[i] = unicode.ToLower(r)
		}
	}
	fmt.Println(string(charList))
}
