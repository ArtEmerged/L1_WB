package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, &wg)

	//work ...
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()

}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("начинаю работу")
	for {
		select {
		case <-ctx.Done():
			log.Println("заканчиваю работу")
			return
		default:
			//work ...
			time.Sleep(time.Second)
		}

	}
}
