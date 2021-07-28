package main

import (
	"fmt"
)

func main() {
	val := 0

	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)

	fmt.Println(fibonacci(val))
}

func fibonacci(val int) []int {
	if val < 2 {
		var slice []int

		return slice
	} else if val == 2 {
		slice := []int{1, 1}

		return slice
	}

	slice := []int{1, 1}
	for i := 2; i < val; i++ {
		slice = append(slice, slice[i-2]+slice[i-1])
	}

	return slice
}

// 模範解答
// func fibonacci(n int) []int {
//     if n < 2 {
//         return make([]int, 0)
//     }

//     nums := make([]int, n)
//     nums[0], nums[1] = 1, 1

//     for i := 2; i < n; i++ {
//         nums[i] = nums[i-1] + nums[i-2]
//     }

//     return nums
// }

// func main() {
//     var num int

//     fmt.Print("What's the Fibonacci sequence you want? ")
//     fmt.Scanln(&num)
//     fmt.Println("The Fibonacci sequence is:", fibonacci(num))
// }
