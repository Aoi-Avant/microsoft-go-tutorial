package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aoi/bank"
)

var accounts = map[float64]*CustomAccount{}

func main() {
	// 初期化
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "john",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number:  1001,
			Balance: 500,
		},
	}
	accounts[1002] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "aoi",
				Address: "Fukuoka, Japan",
				Phone:   "(000) 000 0000",
			},
			Number:  1002,
			Balance: 500,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/remittance", remittance)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 明細
func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Accont with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

// 預金
func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

// 引き出し
func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

// 送金
func remittance(w http.ResponseWriter, req *http.Request) {
	fromqs := req.URL.Query().Get("from")
	toqs := req.URL.Query().Get("to")
	amountqs := req.URL.Query().Get("amount")

	if fromqs == "" {
		fmt.Fprintf(w, "FromAccount number is missing!")
		return
	}

	if toqs == "" {
		fmt.Fprintf(w, "ToAccount number is missing!")
		return
	}

	if from, err := strconv.ParseFloat(fromqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid fromAccount number!")
	} else if to, err := strconv.ParseFloat(toqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid toAccount number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		fromAccount, fromOk := accounts[from]
		toAccount, toOk := accounts[to]
		if !fromOk {
			fmt.Fprintf(w, "FromAccount with number %v can't be found!", from)
		} else if !toOk {
			fmt.Fprintf(w, "ToAccount with number %v can't be found!", to)
		} else {
			err := fromAccount.Remittance(toAccount.Account, amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, fromAccount.Statement())
				fmt.Fprintf(w, toAccount.Statement())
			}
		}
	}
}

// over ride
type CustomAccount struct {
	*bank.Account
}

func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)

	if err != nil {
		return err.Error()
	}

	return string(json)
}
