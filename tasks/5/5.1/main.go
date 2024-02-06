package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func senderWithTimeout(n int) <-chan int {
	ch := make(chan int)

	timeout := time.After(time.Duration(n) * time.Second)

	go func() {
		log.Println("sender запущен")
		for {
			select {
			case ch <- rand.Intn(100):
			case <-timeout:
				log.Println("senderWithTimeout завершился")
				close(ch)
				return
			}
		}

	}()
	return ch
}

func main() {
	fmt.Println("укажите время выполнения программы в секундах")
	var n int
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("вы указали время работы программы меньше 1 секунды")
		return
	}
	out := senderWithTimeout(n)

	log.Println("worker запущен")
	for value := range out {
		fmt.Println(value)
	}
	log.Println("worker завершился")

}
