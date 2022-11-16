package shapes_test

import (
  shapes "github.com/turnerem/shapes"
  "testing"
)

func TestPerimeter(t *testing.T) {
  got := Perimeter(10.0, 10.0)
  want := 40

  if got != want {
    t.Errorf("got %d but want %d", got, want)
  }
}
