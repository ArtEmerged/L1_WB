package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	out := sqrtIn(arr)
	var res int
	for range arr {
		res += <-out
	}
	fmt.Println(res)
}

func sqrtIn(in []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}
