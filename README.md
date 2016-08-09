dice
====
[![GoDoc](https://godoc.org/github.com/mohae/json2go?status.svg)](https://godoc.org/github.com/mohae/dice)[![Build Status](https://travis-ci.org/mohae/dice.png)](https://travis-ci.org/mohae/dice)

Dice is a package for die or dice generation.  Specify the number of sides of a die; for more than one die, specify the dice quantity.

The dice roll is done using `crypto/rand`.  There is a [PCG](http://www.pcg-random.org/) based PRNG implementation under `github.com/mohae/dice/quick`. This is mostly for testability, but may be used if a CSPRNG isn't necessary for your use case.

An example program is in the `dice/roll` directory.

## Usage

    import "github.com/mohae/dice "

    // die w 6 sides
    d := dice.NewDie(6)

    n, err := d.Roll()
    if err != nil {
        // handle error
    }
    fmt.Println(n)

## Roll
Roll is the example cli app. It accepts a `-sides` flag that tells it how many sides the die should have and an optional paramater that specifies how many rolls should be done.

If `-sides` isn't used, the die will default to 6 sides. If the number of rolls isn't passed, the die will only be rolled once. The number of rolls must be a valid number.

If you don't have the repo in your GOPATH:

	go install github.com/mohae/dice/roll

Or in the `roll` directory:

    go install

Run `roll`:
    roll
    5

	roll 3
    4
    1
    5

    roll -sides=20
    14
