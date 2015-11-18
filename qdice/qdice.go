// Package qdice provides non-CSPRNG dice.  Rolling this die is much quicker
// than the CSPRNG based die and does not generate garbage.
//
// The random number generator needs to be seeded.
package qdice

import (
	"math/rand"
)

// Die has n sides.  Use New(n) to get an n-sided die.  This die uses
// math/rand. which is not a CSPRNG.  Don't forget to seed math/rand.
type Die struct {
	sides int
}

// New returns a n-sided die.
func New(n int) Die {
	return Die{n}
}

// Roll returns the result of a die roll. This is not CSPRNG.  Seeding of
// math/rand is required before the first Roll.
func (d Die) Roll() int {
	return rand.Intn(d.sides) + 1
}
