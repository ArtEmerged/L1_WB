package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	sleep(time.Millisecond * 2)
	fmt.Println(time.Since(t))
}

func sleep(n time.Duration) {
	start := time.Now()
	for time.Since(start) < n {
	}
}
