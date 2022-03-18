package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func main() {
	p1 := Point{
		X: 7.0,
		Y: 5.0,
	}

	p2 := Point{
		X: 3.0,
		Y: 2.0,
	}

	fmt.Println(Distance(p1, p2))
	fmt.Println(p1.Distance(p2))
}

func Distance(p, q Point) float64 {
	// Hypot：对于给定的直角三角形的两个直角边，求其斜边的长度
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
