// Roll is an example CLI die app that rolls a die n times where n is a number
// passed to the application as an arg.  By default roll uses a 6 sided die.
// This can be changed using the -sides flag.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mohae/dice"
)

var (
	sides = flag.Int("sides", 6, "number of sides for the die")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  rolldice [-sides] n\n")
	fmt.Fprintf(os.Stderr, "\twhere n is the number of times you want to roll the die\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	args := flag.Args()
	var rolls int
	var err error
	if len(args) > 0 {
		rolls, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not convert %s to an int: %s", args[0], err)
			os.Exit(1)
		}
	}
	// rolls must be > 0
	if rolls == 0 {
		rolls = 1
	}
	d := dice.New(*sides)
	for i := 0; i < rolls; i++ {
		v, err := d.Roll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(v)
	}
	os.Exit(0)
}
