package shapes

import "math"

// Shape - interface to define operations on the possible shapes
type Shape interface {
	Area() float64
}

// Rectangle struct to define properties of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area - calculates area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter - calculates perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle struct to define properties of a Circle
type Circle struct {
	Radius float64
}

// Area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
