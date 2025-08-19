package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * 2
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func printShapeDetails(s [4]Shape) {
	for _, shape := range s {
		fmt.Println(shape.Area())
		fmt.Println(shape.Perimeter())
		switch v := shape.(type) {
		case Circle:
			fmt.Printf("This is a circle with radius %f.\n", v.radius)
		case Rectangle:
			fmt.Printf("This is a rectangle with width %f and height %f.\n", v.width, v.height)
		default:
			fmt.Println("Unknown Shape")
		}
	}
}

func main() {
	c1 := Circle{radius: 5.2}
	c2 := Circle{radius: 3.4}
	r1 := Rectangle{width: 20.2, height: 5.8}
	r2 := Rectangle{width: 10.5, height: 7.2}

	var s = [4]Shape{c1, c2, r1, r2}

	printShapeDetails(s)

}
