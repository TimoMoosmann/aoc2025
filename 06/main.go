package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type MathColumn struct {
	numbers          []uint64
	isMultiplication bool
}

func main() {
    if (len(os.Args) < 2) {
        fmt.Println("Please specify the input file.")
        os.Exit(1)
    }
    inputFileName := os.Args[1]

	part1Result := getResultForMathColumns(
		getMathColumnsPart1(getInputScanner(inputFileName)),
	)
	part2Result := getResultForMathColumns(
		getMathColumnsPart2(getInputScanner(inputFileName)),
	)

	fmt.Printf("Part 1 Result: %v\n", part1Result)
	fmt.Printf("Part 2 Result: %v\n", part2Result)
}

func getMathColumnsPart2(inputScanner *bufio.Scanner) []*MathColumn {
	mathColumnPtrs := make([]*MathColumn, 0)

	lines := make([]string, 0)
	for inputScanner.Scan() {
		lines = append(lines, inputScanner.Text())
	}

	currentMathColumnPtr := &MathColumn{
		numbers: make([]uint64, 0),
	}
	mathColumnPtrs = append(mathColumnPtrs, currentMathColumnPtr)
	for col := 0; col < len(lines[0]); col++ {
		colDigits := 0
		isEmptyCol := true
		var rowValue uint64 = 0
		for row := len(lines) - 1; row >= 0; row-- {
			currentChar := lines[row][col]

			if currentChar == ' ' {
				continue
			} else {
                isEmptyCol = false
            }

			switch currentChar {
			case '*':
				currentMathColumnPtr.isMultiplication = true
			case '+':
				currentMathColumnPtr.isMultiplication = false
			default:
				currentNum := uint64(currentChar - '0')
				rowValue += currentNum * uint64(math.Pow10(colDigits))
				colDigits++
			}
		}
		if isEmptyCol {
			currentMathColumnPtr = &MathColumn{
				numbers: make([]uint64, 0),
			}
			mathColumnPtrs = append(mathColumnPtrs, currentMathColumnPtr)
		} else {
			currentMathColumnPtr.numbers = append(
				currentMathColumnPtr.numbers, rowValue,
			)
		}
	}
	return mathColumnPtrs
}

func getMathColumnsPart1(inputScanner *bufio.Scanner) []*MathColumn {
	mathColumnPtrs := make([]*MathColumn, 0)

    for inputScanner.Scan() {
        line := inputScanner.Text()

        for i, numberStr := range strings.Fields(line) {
            // Create initial column object if none exists for this column
            if len(mathColumnPtrs) < i+1 {
                mathColumnPtrs = append(mathColumnPtrs, &MathColumn{
                    numbers: make([]uint64, 0),
                })
            }

            switch numberStr {
            case "*":
                mathColumnPtrs[i].isMultiplication = true
            case "+":
                mathColumnPtrs[i].isMultiplication = false
            default:
                number, parseErr := strconv.ParseUint(numberStr, 10, 64)
                if parseErr != nil {
                    log.Fatal("Could not parse given number: " + parseErr.Error())
                }

                mathColumnPtrs[i].numbers = append(
                    mathColumnPtrs[i].numbers,
                    number,
                )
            }
        }
    }

	return mathColumnPtrs
}

func getResultForMathColumns(mathColumnPtrs []*MathColumn) uint64 {
	var result uint64 = 0
	for _, mathLine := range mathColumnPtrs {
		var rowResult uint64
		if mathLine.isMultiplication {
			rowResult = 1
		} else {
			rowResult = 0
		}

		for _, number := range mathLine.numbers {
			if mathLine.isMultiplication {
				rowResult *= number
			} else {
				rowResult += number
			}
		}
		result += rowResult
	}
	return result
}

func getInputScanner(inputFileName string) (*bufio.Scanner) {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}

	var inputReader io.Reader = file
	return bufio.NewScanner(inputReader)
}
