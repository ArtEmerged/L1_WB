package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type cache struct {
	data map[int]int
	mu   sync.Mutex
}

func main() {
	newCache := &cache{
		data: make(map[int]int),
	}
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		numW := i + 1
		go newCache.worker(done, &wg, numW)
	}
	// work...
	time.Sleep(time.Second * 5)
	close(done)

	wg.Wait()
	fmt.Println(newCache.data)
}

func (c *cache) worker(done chan struct{}, wg *sync.WaitGroup, numW int) {
	defer wg.Done()
	log.Printf("worker[%d] запущен\n", numW)
	for {
		select {
		case <-done:
			log.Printf("worker[%d] завершился\n", numW)
			return
		default:
			key, value := rand.Intn(100), rand.Intn(100)
			c.mu.Lock()
			c.data[key] = value
			c.mu.Unlock()
			fmt.Printf("worker[%d] записал %d:%d\n", numW, key, value)
			time.Sleep(time.Second)
		}
	}
}
