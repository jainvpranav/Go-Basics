package main

import "fmt"

type triangle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("The area of the shape is: ", s.getArea())
}
