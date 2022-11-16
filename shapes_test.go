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
    t.Run("rectangles area", func(t *testing.T) {

    rect := shapes.Rectangle{10.0, 10.0}

    got := rect.Area()
    want := 100.0

    if got != want {
      t.Errorf("got %.2f but want %.2f", got, want)
    }
  })

  t.Run("circles area", func(t *testing.T) {
    circle := shapes.Circle{10}

    got := circle.Area()
    want := math.Pi * 10 * 10

    if got != want {
      t.Errorf("got %.2f but want %.2f", got, want)
    }
  })
}
