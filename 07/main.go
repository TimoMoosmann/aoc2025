package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
    if (len(os.Args) < 2) {
        fmt.Println("Please specify the input file.")
        os.Exit(1)
    }
    inputFileName := os.Args[1]

    partAAndBWithGrid(inputFileName)
}


func partAAndBWithGrid(inputFileName string) {
    inputScanner, file := getInputScanner(inputFileName)
    defer file.Close()

    // handle first line
    if !inputScanner.Scan() {
        panic("Could not scan first line of input")
    }

    firstLine := inputScanner.Text()
    prevInboundBeamsCount := make([]int, len(firstLine))

    for x, symbol := range firstLine {
        if symbol == 'S' {
            prevInboundBeamsCount[x] = 1
        }
    }

    numSplits := 0
    for inputScanner.Scan() {
        line := inputScanner.Text()
        inboundBeamsCount := make([]int, len(line))
        for x, symbol := range line {
            switch symbol {
            case '^':
                if prevInboundBeamsCount[x] > 0 {
                    numSplits++
                }
                if x - 1 >= 0 {
                    inboundBeamsCount[x - 1] += prevInboundBeamsCount[x]
                }
                if x + 1 < len(line) {
                    inboundBeamsCount[x + 1] += prevInboundBeamsCount[x]
                }
            case '.':
                inboundBeamsCount[x] += prevInboundBeamsCount[x]
            }
        }
        prevInboundBeamsCount = inboundBeamsCount
    }

    numTimelines := 0
    for _, waysToReachBottomPoint := range prevInboundBeamsCount {
        numTimelines += waysToReachBottomPoint
    }

    fmt.Printf("Part A: Beam was split %v times.\n", numSplits)
    fmt.Printf("Part B: %v ways to reach the bottom.\n", numTimelines)
}

func partAWithOutGrid(inputFileName string) {
    inputScanner, file := getInputScanner(inputFileName)
    defer file.Close()

    numSplits := 0
    prevRowBeams := make([]int, 0)

    for inputScanner.Scan() {
        line := inputScanner.Text()
        rowBeams := make([]int, 0)
        for x, symbol := range line {
            switch symbol {
            case 'S':
                rowBeams = append(rowBeams, x)
            case '^':
                if slices.Contains(prevRowBeams, x) {
                    numSplits++
                    if x - 1 >= 0 {
                        rowBeams = append(rowBeams, x - 1)
                    }
                    if x + 1 < len(line) {
                        rowBeams = append(rowBeams, x + 1)
                    }
                }
            case '.':
                if slices.Contains(prevRowBeams, x) {
                    rowBeams = append(rowBeams, x)
                }
            }
        }
        prevRowBeams = rowBeams
    }

    fmt.Printf("Part A: Beam was split %v times", numSplits)
}

func getInputScanner(inputFileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}

	var inputReader io.Reader = file
	return bufio.NewScanner(inputReader), file
}
