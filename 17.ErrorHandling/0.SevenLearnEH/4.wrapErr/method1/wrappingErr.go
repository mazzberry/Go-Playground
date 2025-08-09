package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := CopyFile("src.txt", "dest.txt")
	if err != nil {
		fmt.Println("error: %v\n", err)
	}
}

func CopyFile(sourceName, destinationName string) (err error) {

	source, err := os.Open(sourceName)

	if err != nil {
		return fmt.Errorf("during copy file could not open the source file: %w", err)
	}

	//defer source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return fmt.Errorf("during copy file could not open the source file: %w", err)
	}

	
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	//defer destination.Close()

	return nil

}
