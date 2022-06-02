package main

import (
	"fmt"
	"math"
)

const Pi float64 = math.Pi

type Rectangle struct {
	name    string
	sideLen float64
}

func NewRectangle() Rectangle {
	return Rectangle{
		"прямоугольник",
		5,
	}
}

type Circle struct {
	name    string
	sideLen float64
}

func NewCircle() Circle {
	return Circle{
		"окружность",
		5,
	}
}

func (c Circle) Area() float64 {
	return Pi * math.Pow(c.sideLen, 2)
}

func (c Circle) Type() string {
	return c.name
}

func (r Rectangle) Area() float64 {
	return r.sideLen * r.sideLen
}

func (r Rectangle) Type() string {
	return r.name
}

func main() {
	circle := NewCircle()
	rectangle := NewRectangle()

	fmt.Printf("Площадь %s = %f\n ", circle.Type(), circle.Area())
	fmt.Printf("Площадь %s = %f\n ", rectangle.Type(), rectangle.Area())
}
