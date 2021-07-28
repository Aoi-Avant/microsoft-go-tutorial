package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fib(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()

	command := ""
	data := make(chan int)

	go fib(data)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fib(ch chan float64, number float64) {
// 	x, y := 1.0, 1.0
// 	for i := 0; i < int(number); i++ {
// 		x, y = y, x+y
// 	}

// 	ch <- x
// }

// func main() {
// 	start := time.Now()

// 	ch1 := make(chan string)
// 	ch2 := make(chan float64)
// 	i := 1
// 	var s string

// 	fmt.Println(i)

// 	for {
// 		go inputScan(ch1, s)

// 		select {
// 		case scan := <-ch1:
// 			if scan != "quit" {
// 				go fib(ch2, float64(i))
// 				fmt.Printf("%v\n", <-ch2)
// 				i++
// 			} else {
// 				fmt.Println("Done calculating Fibonacci!")
// 				elapsed := time.Since(start)
// 				fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
// 				break
// 			}
// 		}
// 	}
// }

// func inputScan(ch chan string, s string) {
// 	fmt.Scanf("%s", &s)

// 	ch <- s
// }
