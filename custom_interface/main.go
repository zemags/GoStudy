package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{
		height: 10.0,
		base:   10.0,
	}

	sq := square{
		sideLength: 10.0,
	}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) float64 {
	result := s.getArea()
	fmt.Println(result)
	return result
}
