package main

import "fmt"

type CheckinAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CheckinAccount) ToWithdraw(withdrawalAmount float64) string {
	canWithdraw := withdrawalAmount >= 0 && withdrawalAmount <= c.balance
	if canWithdraw {
		c.balance -= withdrawalAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckinAccount) Deposit(amountDeposited float64) (string, float64) {
	if amountDeposited > 0 {
		c.balance += amountDeposited
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Valor do depósito menor que zero", c.balance
	}
}

func main() {
	accoount01 := CheckinAccount{}
	accoount01.holder = "Cliente01"
	accoount01.balance = 5000

	fmt.Println(accoount01.ToWithdraw(6000))
	fmt.Println(accoount01.Deposit(200))

}
