package main

import (
	"bank/accounts"
	"fmt"
)

func PayBill(account checkAccount, billetValue float64) {
	account.ToWithdraw(billetValue)
}

type checkAccount interface {
	ToWithdraw(value float64) string
}

func main() {

	// accountClien01 := accounts.SavingsAccount{}
	// accountClien01.Deposit(-100)

	// fmt.Println(accountClien01.GetBalance())

	// accounClient02 := accounts.SavingsAccount{}
	// accounClient02.Deposit(-100)

	accounClient03 := accounts.CheckinAccount{}
	accounClient03.Deposit(100)
	PayBill(&accounClient03, 60)

	fmt.Println(accounClient03.GetBalance())
}
