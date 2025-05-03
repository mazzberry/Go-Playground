package main // 3.2.1

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
	go start2()
	fmt.Println("In goroutine")
}

func start2() {
	fmt.Println("In goroutine 2")

}