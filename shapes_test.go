package shapes_test

import (
  shapes "github.com/turnerem/shapes"
  "testing"
  "math"
)

func TestPerimeter(t *testing.T) {
  t.Run("rectangles perim", func(t *testing.T) {

    rect := shapes.Rectangle{10.0, 10.0}

    got := rect.Perimeter()
    want := 40.0

    if got != want {
      t.Errorf("got %.2f but want %.2f", got, want)
    }
  })

  t.Run("circles perim", func(t *testing.T) {
    circle := shapes.Circle{10.0}

    got := circle.Perimeter()
    want := 2 * math.Pi * 10.0

    if got != want {
      t.Errorf("got %.2f but want %.2f", got, want)
    }
  })
}

func TestArea(t *testing.T) {

  areaTests := []struct {
    name	string
    shape	shapes.Shape
    want	float64
  }{
    {name: "Rectangle", shape: shapes.Rectangle{10.0, 10.0}, want: 100.0},
    {name: "Circle", shape: shapes.Circle{10.0}, want: math.Pi * 100.0},
    {name: "Triangle", shape: shapes.Triangle{10.0, 10.0}, want: 50.0},
  }

  for _, tt := range areaTests {
    t.Run(tt.name, func(t *testing.T) {

      got := tt.shape.Area()

      if got != tt.want {
        t.Errorf("got %.2f but want %.2f", got, tt.want)
      }
    })
  }
}
