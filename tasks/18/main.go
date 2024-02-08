package main

import (
	"fmt"
	"sync"
)

type counter struct {
	num int
	mu  sync.Mutex
}

func (c *counter) intrepret(wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.mu.Unlock()
	c.mu.Lock()
	c.num++

}

func main() {
	counter := counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.intrepret(&wg)
	}
	wg.Wait()
	fmt.Println(counter.num)
}
