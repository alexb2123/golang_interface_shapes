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
	sa := square{2}
	ta := triangle{2, 2}

	printArea(sa)
	printArea(ta)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {

	area := .5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
