package main

import (
	"os"
	"strconv"
)

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	println("Sum:", sum)
	println("Mul:", mul)

	sum2, _ := calc(os.Args[1], os.Args[2])
	println("Sum:", sum2)

	firstName := "John"
	updateName(&firstName)
	println(firstName)
}

// カスタム関数
func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum := int1 + int2
	return sum
}

func calc(number1 string, number2 string) (int, int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum := int1 + int2
	mul := int1 * int2
	return sum, mul
}

// ポインター
func updateName(name *string) {
	*name = "David"
}
