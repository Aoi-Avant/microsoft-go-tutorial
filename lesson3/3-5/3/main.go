package main

import "fmt"

func main() {
	for {
		val := 0

		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered: ", val)
		}
	}
}
