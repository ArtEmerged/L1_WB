package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(done, &wg)

	//work ...
	time.Sleep(time.Second * 5)
	close(done)
	wg.Wait()

}

func worker(done chan struct{}, wg *sync.WaitGroup) {
	log.Println("начинаю работу")
	for {
		select {
		case <-done:
			log.Println("заканчиваю работу")
			wg.Done()
			return
		default:
			//work ...
			time.Sleep(time.Second)
		}

	}
}
