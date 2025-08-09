package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type IOError struct {
	FileName string
	Message  string
	Err      error
}

func (error *IOError) Unwrap() error {
	return error.Err
}

func (error IOError) Error() string {
	return fmt.Sprintf("IO error occurred... file name %s message %s Detail: %s", error.FileName, error.Message, error.Err.Error())
}

func main() {
	err := CopyFile("src.txt", "dest.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		fmt.Printf("Unwrapped error: %s", errors.Unwrap(err))
	}
}

func CopyFile(sourceName, destinationName string) (err error) {

	source, err := os.Open(sourceName)

	if err != nil {
		return &IOError{FileName: sourceName, Message: "during copy file could not open source file", Err: err}
	}

	//defer source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return &IOError{FileName: sourceName, Message: "during copy file could not create destination file", Err: err}
	}

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	//defer destination.Close()

	return nil

}
