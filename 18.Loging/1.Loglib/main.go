package main

import (
	"fmt"
	"log"
	"os"
)

var (
	errorLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
)

func init() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("something went wrong with log file", err)
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile

	log.SetFlags(flags)
	log.SetOutput(file)

	errorLogger = log.New(file, "Error : ", flags)
	warnLogger = log.New(file, "warning : ", flags)
	infoLogger = log.New(file, "info : ", flags)
}

func main() {
	warnLogger.Println("start of main func\n")
	sum(2, 3)
	fmt.Println("hello")
	warnLogger.Println("\nend main func\n")

}

func sum(a, b int) {
	infoLogger.Println("start of sum\n")
	fmt.Println(a + b)
	errorLogger.Println("\nend of the sum operation\n")
}
