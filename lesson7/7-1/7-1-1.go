package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api)
	}

	time.Sleep(3 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seccond!\n", elapsed.Seconds())
}

func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error: %s is down!\n", api)
		return
	}

	fmt.Printf("Success: %s is up and running!\n", api)
}
