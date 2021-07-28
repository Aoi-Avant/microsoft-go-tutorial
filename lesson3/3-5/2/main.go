package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	x, _ := strconv.ParseFloat(input, 64)

	var sqroot float64 = 1
	for i := 1; i <= 10; i++ {
		old := sqroot
		sqroot = old - (old*old-x)/(2*old)

		if old == sqroot {
			fmt.Println("Square root is: ", sqroot)
			break
		}

		fmt.Println("A guess for square root is ", sqroot)
	}
}
