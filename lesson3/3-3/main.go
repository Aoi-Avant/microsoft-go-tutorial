package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1...100 is", sum)

	// 空の前ステートメントと後ステートメント
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	// 無限ループ
	var num2 int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writting inside the loop ...")
		if num2 = rand.Int31n(10); num2 == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num2)
	}

	// continue
	sum2 := 0
	for num3 := 1; num3 <= 100; num3++ {
		if num3%5 == 0 {
			continue
		}
		sum2 += num3
	}

	fmt.Println("the sum of 1 to 100, but excluding numbers divisible by 5, is ", sum2)
}
