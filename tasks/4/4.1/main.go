package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var N int
	fmt.Println("укажите количество воркеров: 0 < N < 100")
	fmt.Scan(&N)

	if N < 1 || N > 99 {
		fmt.Printf("вы указали %d воркеров. количество воркеров: 0 < N < 100 ", N)
		return
	}
	ch := make(chan int)
	var wg sync.WaitGroup

	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for i := 0; i < N; i++ {
		wg.Add(1)
		i := i + 1
		go worker(quit, &wg, ch, i)
	}

	log.Println("sender запущен")
	for {
		select {
		case <-quit.Done():
			close(ch)
			log.Println("sender завершился")
			wg.Wait()
			return
		default:
			time.Sleep(time.Second)
			numSend := rand.Intn(100)
			ch <- numSend
			log.Printf("sender передал: %d\n", numSend)
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, numW int) {
	log.Printf("worker[%d] запущен\n", numW)
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker[%d] завершился\n", numW)
			wg.Done()
			return
		default:
			out, ok := <-ch
			if ok {
				fmt.Printf("worker[%d] выводит %d\n", numW, out)
			}
		}
	}
}
