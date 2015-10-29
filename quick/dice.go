// Package quick provides a non-CSPRNG die.  Rolling this die is much quicker
// than the CSPRNG based die and does not generate garbage.  Odds are the
// crypto/rand based die should be used.
package quick

import (
	"math/rand"
)

// Die has n sides.  Use New(n) to get an n-sided die.  This die uses
// math/rand. which is not a CSPRNG.  Don't forget to seed math/rand.
type Die struct {
	sides int
}

// New returns a n-sided die.
func New(i int) Die {
	return Die{i}
}

// Roll returns the result of a die roll. This is not CSPRNG.  Seeding of
// math/rand is required before the first Roll.
func (d Die) Roll() int {
	return rand.Intn(d.sides) + 1
}
