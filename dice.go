package dice

import (
  "crypto/rand"
  "math/big"
)

type Die struct {
  sides int64
}

// New returns a Die with i sides.
func New(i int) Die {
  return Die{int64(i)}
}

// Roll rolls the die. It returns either the roll results or an error.
func (d Die) Roll() (int, error) {
  b := big.NewInt(d.sides + 1)
  r, err := rand.Int(rand.Reader, b)
  if err != nil {
    return 0, err
  }
  return int(r.Int64()), nil
}
