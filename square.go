package main

import (
	"fmt"
	"log"
	"math"
)

// Point ia a 2D point
type Point struct {
	X float64
	Y float64
}

func (p *Point) Move(dx float64, dy float64) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length float64
}

func NewSquare(x float64, y float64, length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length must be > 0 (was %f)", length)
	}

	square := &Square{
		Center: Point{X: x, Y: y},
		Length: length,
	}

	return square, nil
}

func (s *Square) Move(dx float64, dy float64) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s, err := NewSquare(3, 3, 4)
	if err != nil {
		//fmt.Printf("Error: can't create a square - %s\n", err)
		//os.Exit(1)
		log.Fatalf("Error: can't create a square - %s\n", err)

	}
	s.Move(2, 3)

	fmt.Printf("Area of the square %+v is %f\n", s, s.Area())

	c := &Circle{10}
	fmt.Printf("Area of the circle %+v is %f\n", c, c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Printf("The sum area of the square and the circle is %f\n", sa)
}
