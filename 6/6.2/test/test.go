package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p1 := &Point{1, 2}
	p1.Add1()
	fmt.Println(*p1)

	p2 := Point{1, 2}
	p2.Add2()
	fmt.Println(p2)

	Point{1, 2}.Add2()
	(&Point{1, 2}).Add1()

}

func (p *Point) Add1() {
	p.X++
	p.Y++
}

func (p Point) Add2() {
	p.X++
	p.Y++
}
