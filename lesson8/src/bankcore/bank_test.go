package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("Accountが正常に作成されていません")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("預金処理が正常に動いていません")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("預金額は0以上である必要があります")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("引き出しの処理が正常に行われていません")
	}
}

func TestWithdrawInvalidNegative(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Withdraw(0); err == nil {
		t.Error("引き出す金額は0より大きい必要があります")
	}
}

func TestWithdrawInvalidLessThan(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(1)

	if err := account.Withdraw(2); err == nil {
		t.Error("預金額を超えて引き出すことはできません")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()

	if statement != "1001 - John - 100" {
		t.Error("明細のフォーマットが正しくありません")
	}
}

func TestRemittance(t *testing.T) {
	fromAccount := Account{
		Customer: Customer{
			Name:    "john",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 500,
	}
	toAccount := Account{
		Customer: Customer{
			Name:    "aoi",
			Address: "Fukuoka, Japan",
			Phone:   "(000) 000 0000",
		},
		Number:  1002,
		Balance: 500,
	}

	err := fromAccount.Remittance(&toAccount, 100)

	if err != nil {
		t.Error("総金額が正常ではありません")
	}

	fromStatement := fromAccount.Statement()
	toStatement := toAccount.Statement()

	if fromStatement != "1001 - john - 600" || toStatement != "1002 - aoi - 400" {
		t.Error("送金が正常に行われていません")
	}
}
