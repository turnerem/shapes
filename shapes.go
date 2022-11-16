package shapes

import (
  "math"
)

type Shape interface {
  Perimeter()	float64
  Area()	float64
}

type Rectangle struct {
  Width		float64
  Height	float64
}

type Circle struct {
  Radius	float64
}

type Triangle struct {
  Width		float64
  Height	float64
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.Radius
}

func (t Triangle) Perimeter() float64 {
  return 0.5 * t.Width * t.Height
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
  return 0.5 * t.Width * t.Height
}
