package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{q.X + p.X, q.Y + p.Y} }
func (p Point) Sub(q Point) Point { return Point{q.X - p.X, q.Y - p.Y} }

type Path []Point

func main() {
	point := Point{1, 2}
	point2 := Point{1, 2}
	path := make(Path, 0, 1)
	path = append(path, point2)
	path.TranslateBy(point, true)
	fmt.Println(path[0])
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	// o := Point.Add
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}
