package main

import (
	"errors"
	"fmt"
)

type Account struct {
	LastName  string
	FirstName string
}

func (account *Account) ChangeName(lastName string, firstName string) {
	account.LastName = lastName
	account.FirstName = firstName
}

type Employee struct {
	Account
	Credit float64
}

func (employee Employee) String() string {
	return fmt.Sprintf("姓は%s、名は%s、この看板背負って始まった%.2f円", employee.LastName, employee.FirstName, employee.Credit)
}

func CreateEmployee(lastName string, firstName string, credit float64) (*Employee, error) {
	return &Employee{
		Account{lastName, firstName},
		credit,
	}, nil
}

func (employee *Employee) AddCredits(credit float64) (float64, error) {
	if credit > 0.0 {
		employee.Credit += credit
		return employee.Credit, nil
	}

	return 0.0, errors.New("Creditを足せません")
}

func (employee *Employee) RemoveCredits(credit float64) (float64, error) {
	if employee.Credit > 0.0 {
		if credit > 0.0 {
			employee.Credit -= credit
			return employee.Credit, nil
		}
		return 0.0, errors.New("Creditを引けません")
	}
	return 0.0, errors.New("Creditから引けません")
}

func (employee *Employee) CheckCredits() float64 {
	return employee.Credit
}

func main() {
	employee, _ := CreateEmployee("John", "Doe", 1000)
	fmt.Println(employee)
	employee.ChangeName("IceBahn", "Fork")
	fmt.Println(employee)
	fmt.Println(employee.CheckCredits())

	credits, err := employee.AddCredits(1500)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("現在の残額: %.2f円\n", credits)
	}

	_, err = employee.RemoveCredits(2000)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(employee.CheckCredits())
}
