package shapes_test

import (
  shapes "github.com/turnerem/shapes"
  "testing"
)

func TestPerimeter(t *testing.T) {
  rect := shapes.Rectangle{10.0, 10.0}

  got := shapes.Perimeter(rect)
  want := 40.0

  if got != want {
    t.Errorf("got %.2f but want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  rect := shapes.Rectangle{10.0, 10.0}

  got := shapes.Area(rect)
  want := 100.0

  if got != want {
    t.Errorf("got %.2f but want %.2f", got, want)
  }
}
