package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// defer
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	// ファイルを閉じる
	f, err := os.Create("notes.txt")
	if err != nil {
		return
	}
	defer f.Close()

	if _, err = io.WriteString(f, "Learning Go!"); err != nil {
		return
	}

	f.Sync()

	// recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// panic
	g(0)
	fmt.Println("Program finished successfully!")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicling!")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()")
	g(i + 1)
}
