package main

import (
	"fmt"
	"math"
)

type intVar []int8

func main() {

	x := newIntVar()

	x.printEvenOrOdd()

}

func newIntVar() intVar {

	y := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return y
}

func (i intVar) printEvenOrOdd() {

	for _, num := range i {

		if math.Mod(float64(num), 2) == 0 {
			fmt.Println(num, " is Even")
		} else {

			fmt.Println(num, " is Odd")
		}

	}

}
