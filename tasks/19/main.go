package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("incorrect input")
		return
	}

	word := []rune(arg[0])
	if len(word) == 0 {
		return
	}
	if len(word) > 1 {
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			word[i], word[j] = word[j], word[i]
		}
	}
	fmt.Printf("%s - %s\n", arg[0], string(word))
}
