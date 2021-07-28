package main

import "fmt"

func main() {
	x := 27

	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	// 複合ifステートメント
	if num := giveMeNumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// ifの外では使えない
	// fmt.Println(num)
}

func giveMeNumber() int {
	return -1
}
