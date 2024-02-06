package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(wg *sync.WaitGroup, ch <-chan int, numW int) {
	defer wg.Done()
	log.Printf("worker[%d] запущен\n", numW)
	for {
		out, ok := <-ch
		if !ok {
			log.Printf("worker[%d] завершился\n", numW)
			return
		}
		fmt.Printf("worker[%d] выводит %d\n", numW, out)
	}
}

func main() {
	var N int
	fmt.Println("Укажите количество воркеров от 1 до 99")
	fmt.Scan(&N)

	if N < 1 || N > 99 {
		fmt.Printf("Вы указали %d воркеров. Количество воркеров должно быть в диапазоне 1-99\n", N)
		return
	}
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		i := i + 1
		go worker(&wg, ch, i)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	log.Println("sender запущен")
	for {
		select {
		case <-exit:
			close(ch)
			log.Println("sender завершился")
			wg.Wait()
			return
		default:
			time.Sleep(time.Second)
			numSend := rand.Intn(100)
			ch <- numSend
			fmt.Printf("sender передал: %d\n", numSend)
		}
	}
}
