package main

import (
	"fmt"
	"time"
)


func execute(id int) {
	fmt.Printf("id: %d\n", id)
}

func main() {
	fmt.Println("main started")
	for i:=1; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("main finished")
}

