package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var done bool
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)
	go worker(&done, &wg)

	//work ...
	time.Sleep(time.Second * 5)

	mu.Lock()
	done = true
	mu.Unlock()

	wg.Wait()
}

func worker(done *bool, wg *sync.WaitGroup) {
	log.Println("начинаю работу")
	for {
		if *done {
			log.Println("заканчиваю работу")
			wg.Done()
			return
		}
		//work ...
		time.Sleep(time.Second)
	}

}
