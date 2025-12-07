package main

import (
    "bufio"
	"fmt"
    "io"
	"os"
)

func main() {
    if (len(os.Args) < 2) {
        fmt.Println("Please specify the input file.")
        os.Exit(1)
    }
    inputFileName := os.Args[1]

    inputScanner, file := getInputScanner(inputFileName)
    defer file.Close()

    for inputScanner.Scan() {
        inputScanner.Text()
    }
}

func getInputScanner(inputFileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}

	var inputReader io.Reader = file
	return bufio.NewScanner(inputReader), file
}
