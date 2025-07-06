package interfaces

import "math"

// A type implements an interface by implementing its methods. 
// There is no explicit declaration of intent, no "implements" keyword.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// A method is a function with a receiver
// when you call a function or a method the arguments are copied. 
// When calling func (r Rectangle) Area(), the r is a copy of the rectangle
// we called the method from.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}