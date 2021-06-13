package structs

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// method for Triangle type
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Rectangle struct {
	Width  float64
	Height float64
}

// method for Rectangle type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// method to calculate area for Circle type
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
