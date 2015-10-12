package quick

import (
  "math/rand"
  "testing"
)

func TestDie(t *testing.T) {
  results := []int{6, 6, 3, 1, 2, 2, 4, 3}
  rand.Seed(42)
  d := New(6)
  for i, v := range results {
      n := d.Roll()
      if n != v {
        t.Errorf("%d: expected %d, got %d", i, v, n)
      }
  }
}
