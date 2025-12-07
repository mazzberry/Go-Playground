package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func main() {
	mu := &sync.Mutex{}

	go increment(mu)
	go increment(mu)
	go increment(mu)
	go increment(mu)

	time.Sleep(time.Second)
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println("incrementing count :", count)
}
