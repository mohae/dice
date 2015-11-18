// Package dice implements a die that uses crypto/rand for its roll.
package dice

import (
	"crypto/rand"
	"math/big"
)

// Die has n sides. Use New(n) to obtain a n-sided Die.
type Die struct {
	sides int64
}

// New returns a Die with n sides.
func New(n int) Die {
	return Die{int64(n)}
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
