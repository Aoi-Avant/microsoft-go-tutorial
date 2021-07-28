package main

import (
	"fmt"
)

func main() {
	// 配列
	var a [3]int

	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	// 配列の初期化
	cities := [5]string{"NewYork", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	// 配列の省略記号
	cities2 := [...]string{"NewYork", "Paris", "Belrin", "Madrid"}
	fmt.Println("Cities:", cities2)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position", numbers[99])
	fmt.Println("Length:", len(numbers))

	// 多次元配列
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)
}
