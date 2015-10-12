package quick

import (
  "math/rand"
)

type Die struct {
  sides int
}

func New(i int) Die {
  return Die{i}
}

func (d Die) Roll() int {
  return rand.Intn(d.sides) + 1
}
