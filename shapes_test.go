package shapes_test

import (
  shapes "github.com/turnerem/shapes"
  "testing"
)

func TestPerimeter(t *testing.T) {
  got := shapes.Perimeter(10.0, 10.0)
  want := 40.0

  if got != want {
    t.Errorf("got %.2f but want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  got := shapes.Area(10.0, 10.0)
  want := 100.0

  if got != want {
    t.Errorf("got %.2f but want %.2f", got, want)
  }
}
