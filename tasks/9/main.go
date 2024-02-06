package main

import (
	"fmt"
	"sync"
)


func woker(input chan int) <-chan int {
	output := make(chan int)
	go func() {
		for value := range input {
			output <- value * 2
		}
		close(output)
		fmt.Println("workr end")
	}()
	return output
}

func writer(ouput <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range ouput {
		fmt.Println(result)
	}
	fmt.Println("writer end")
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 15, 16, 100, 20, 1, 23, 41, 41, 4, 11, 1, 41, 0}

	var wg sync.WaitGroup
	wg.Add(1)
	input := make(chan int)
	ouput := woker(input)
	go writer(ouput, &wg)
	for _, num := range arr {
		input <- num
	}
	fmt.Println("sender end")
	close(input)
	wg.Wait()
}
