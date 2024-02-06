package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("введите слово:")
	var input string
	fmt.Scan(&input)
	onlyLow := strings.ToLower(input)
	counter := make(map[rune]struct{})
	for _, v := range onlyLow {
		if _, ok := counter[v]; ok {
			fmt.Printf("%s - %v\n", input, false)
			return
		}
		counter[v] = struct{}{}
	}
	fmt.Printf("%s - %v\n", input, true)
}
