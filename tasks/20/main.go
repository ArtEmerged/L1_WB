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

	words := string(arg[0])
	if len(words) == 0 {
		return
	}
	var newWords string
	if len(words) > 1 {
		words += " "
		var begW int

		for i := range words {
			if words[i] == ' ' {
				newWords = " " + words[begW:i] + newWords
				begW = i + 1
			}
		}
	}
	fmt.Printf("|%s| - |%s|\n", arg[0], newWords[1:])
}
