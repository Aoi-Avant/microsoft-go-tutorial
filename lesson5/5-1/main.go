package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInformation(1000)
	if err != nil {
		// エラー処理
		if errors.Is(err, ErrNotFound) {
			fmt.Printf("NOT FOUND: %v\n", err)
		} else {
			fmt.Printf("SOME ERROR: %v\n", err)
		}
	} else {
		fmt.Print(employee)
	}
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	// エラーを伝達
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, err

	// リトライ
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}
		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}
	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
