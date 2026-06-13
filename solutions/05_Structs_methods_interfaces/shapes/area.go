package shapes

import "math"

type Shape interface {
	Area() float64
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
	return (triangle.Base * triangle.Height) * 0.5
}
