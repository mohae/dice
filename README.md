dice
====

Dice is a package for die or dice generation.  Specify the number of sides of a die; for more than one die, specify the dice quantity.

The dice roll is done using `crypto/rand`.  There is a `math/rand` implementation under `github.com/mohae/dice/quick`. This is mostly for testability, but may be used if a CSPRNG isn't necessary for your use case.

## Usage

    import "github.com/mohae/dice "

    // die w 6 sides
    d := dice.NewDie(6)

    n := d.Roll()


