package main

import "math"

type Pos struct {
    x int
    y int
    z int
}

type Connection struct {
    pos1 Pos
    pos2 Pos
    dist float64
}

func EuclidDist(pos1 Pos, pos2 Pos) float64 {
    return math.Sqrt(
        math.Pow(float64(pos1.x) - float64(pos2.x), 2) +
        math.Pow(float64(pos1.y) - float64(pos2.y), 2) +
        math.Pow(float64(pos1.z) - float64(pos2.z), 2))
}

