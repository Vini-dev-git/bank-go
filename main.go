package main

import (
	"bank/accounts"
	"fmt"
)

func main() {

	accountexample := accounts.CheckinAccount{}
	accountexample.Deposit(100)

	fmt.Println(accountexample.GetBalance())

}
