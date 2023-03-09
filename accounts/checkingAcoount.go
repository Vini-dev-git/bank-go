package accounts

import "bank/custumers"

type CheckinAccount struct {
	Holder        custumers.Holder
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckinAccount) ToWithdraw(withdrawalAmount float64) string {
	canWithdraw := withdrawalAmount >= 0 && withdrawalAmount <= c.Balance
	if canWithdraw {
		c.Balance -= withdrawalAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckinAccount) Deposit(amountDeposited float64) (string, float64) {
	if amountDeposited > 0 {
		c.Balance += amountDeposited
		return "Depósito realizado com sucesso", c.Balance
	} else {
		return "Valor do depósito menor que zero", c.Balance
	}
}

func (c *CheckinAccount) Transfer(transferAmount float64, destinationAccount *CheckinAccount) string {
	if transferAmount <= c.Balance && transferAmount > 0 {
		c.Balance -= transferAmount
		destinationAccount.Deposit(transferAmount)
		return "Tranferência feita com sucesso"
	} else {
		return "Tranferência não concluída. Por favor verifique os valores"
	}
}
