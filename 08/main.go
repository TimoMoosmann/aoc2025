package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"maps"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Circuit map[Pos]int

type OrderableCircuits []*Circuit

func (c OrderableCircuits) Len() int {
    return len(c)
}

func (c OrderableCircuits) Less(i, j int) bool {
    return len(*c[i]) > len(*c[j])
}

func (c OrderableCircuits) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}

func main() {
    fmt.Println("AoC 2025 - Day 8")
    if len(os.Args) < 3 {
        fmt.Println(
            "Usage: main.go {input file} {number of connections for part 1}")
        os.Exit(1)
    }
    inputFileName := os.Args[1]
    numberOfConnectionsForPart1 := atoiOrPanic(os.Args[2])

    inputScanner, file := getInputScanner(inputFileName)
    defer file.Close()

    positions := parsePositions(inputScanner)
    resultPart1 := solvePart1(numberOfConnectionsForPart1, positions)
    resultPart2 := solvePart2(positions)

    fmt.Printf("Result for part 1: %d\n", resultPart1)
    fmt.Printf("Result for part 2: %d\n", resultPart2)
}

func solvePart1(numberOfConnections int, positions []Pos) int {
    pq := getConnectionsPriorityQueue(positions)
    ordrableCircuits, posToCircuit := getInitialCircuitsForPositions(positions)

    for range numberOfConnections {
        connection := heap.Pop(&pq).(*Connection)
        mergeCircuits(*connection, posToCircuit)
    }
    sort.Sort(ordrableCircuits)

    return len(*ordrableCircuits[0]) * len(*ordrableCircuits[1]) * len(*ordrableCircuits[2])
}

func solvePart2(positions []Pos) int {
    pq := getConnectionsPriorityQueue(positions)
    _, posToCircuit := getInitialCircuitsForPositions(positions)

    for {
        connection := heap.Pop(&pq).(*Connection)
        mergedCircuit := mergeCircuits(*connection, posToCircuit)

        if len(*mergedCircuit) >= len(positions) {
            return connection.pos1.x * connection.pos2.x
        }
    }
}

func parsePositions(inputScanner *bufio.Scanner) []Pos {
    positions := make([]Pos, 0)
    for inputScanner.Scan() {
        line := inputScanner.Text()
        posStrings := strings.Split(line, ",")

        positions = append(positions, Pos{
            x: atoiOrPanic(posStrings[0]),
            y: atoiOrPanic(posStrings[1]),
            z: atoiOrPanic(posStrings[2]),
            })
    }

    return positions
}

func getConnectionsPriorityQueue(positions []Pos) ConnectionsPriorityQueue {
    pq := make(ConnectionsPriorityQueue, 0)

    for pos1Idx := range positions {
        for pos2Idx := pos1Idx + 1; pos2Idx < len(positions); pos2Idx++ {
            pos1 := positions[pos1Idx]
            pos2 := positions[pos2Idx]
            
            dist := EuclidDist(pos1, pos2)

            heap.Push(&pq, &Connection{pos1, pos2, dist})
        }
    }

    return pq
}

func getInitialCircuitsForPositions(positions []Pos) (
    circuits OrderableCircuits,
    posToCircuit map[Pos]*Circuit,
) {
    circuits = make(OrderableCircuits, 0)
    posToCircuit = make(map[Pos]*Circuit)
    
    for _, pos := range positions {
        singlePosCircuit := make(Circuit)
        singlePosCircuit[pos] = 1
        circuits = append(circuits, &singlePosCircuit)
        posToCircuit[pos] = &singlePosCircuit
    }

    return circuits, posToCircuit
}


func mergeCircuits(
    connection Connection,
    posToCircuit map[Pos]*Circuit,
) *Circuit {
    circuit1 := posToCircuit[connection.pos1]
    circuit2 := posToCircuit[connection.pos2]

    if circuit1 == circuit2 {
        return circuit2
    }

    maps.Copy((*circuit2), *circuit1)
    emptyCircuit := make(Circuit)
    *circuit1 = emptyCircuit

    for pos := range *circuit2 {
        posToCircuit[pos] = circuit2
    }

    return circuit2
}

func atoiOrPanic(numToConv string) int {
    num, err := strconv.Atoi(numToConv)

    if err != nil {
        panic(err)
    }

    return num
}

func getInputScanner(inputFileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}


	var inputReader io.Reader = file
	return bufio.NewScanner(inputReader), file
}

