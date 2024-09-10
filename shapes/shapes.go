package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Round(2*math.Pi*c.Radius*100) / 100
}

func (c Circle) Area() float64 {
	return math.Round(math.Pi*math.Pow(c.Radius, 2.0)*100) / 100
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
