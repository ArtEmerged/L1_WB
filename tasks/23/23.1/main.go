package main

import "fmt"

func main() {
	newSlice := []int{1, 2, 3, 4, 5}
	newSlice = deleteElem(newSlice, 1)
	fmt.Println(newSlice)
}

func deleteElem(old []int, i int) []int {
	if i < 0 || i > len(old)-1 {
		return old
	}
	return append(old[:i], old[i+1:]...)
}
