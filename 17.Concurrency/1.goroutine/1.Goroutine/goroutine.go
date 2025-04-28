package main 

import (
	"fmt"
	"time"
)

func main() {
	go start()
	fmt.Println("started")
	time.Sleep(1 * time.Second)
	fmt.Println("finished")
}

func start() {
	fmt.Println("In goroutine")
}

