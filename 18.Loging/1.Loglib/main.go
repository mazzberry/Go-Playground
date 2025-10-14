package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("something went wrong with log file", err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(file)
}

func main() {
	log.Println("start of main func\n")
	sum(2, 3)
	log.Println("\nend main func\n")

}

func sum(a, b int) {
	log.Println("start of sum\n")
	fmt.Println(a + b)
	log.Println("\nend of the sum operation\n")
}
