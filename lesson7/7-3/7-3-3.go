package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	send2(ch, "Hello World!")
	read(ch)
}

// 送信受信の指定
func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
	// ch <- "Bye!" 受信専用のため送信はできない
}
