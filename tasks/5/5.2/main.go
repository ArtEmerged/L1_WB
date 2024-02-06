package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("укажите время выполнения программы в секундах")

	var t time.Duration
	fmt.Scan(&t)

	t *= time.Second

	if t < time.Second {
		fmt.Println("вы указали время работы программы меньше 1 секунды")
		return
	}
	ctx, stop := context.WithTimeout(context.Background(), t)
	defer stop()

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg, ch)

	log.Println("sender запущен")
	for {
		select {
		case <-ctx.Done():
			close(ch)
			log.Println("sender завершился")
			wg.Wait()
			return
		default:
			numSend := rand.Intn(100)
			ch <- numSend
			log.Printf("sender передал: %d\n", numSend)
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	log.Println("worker[1] запущен")
	for {
		select {
		case <-ctx.Done():
			log.Println("worker[1] завершился")
			return
		default:
			out, ok := <-ch
			if ok {
				fmt.Printf("worker[1] выводит %d\n", out)
			}
		}
	}
}
