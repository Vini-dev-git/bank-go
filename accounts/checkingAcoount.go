package accounts

import "bank/custumers"

type CheckinAccount struct {
	Holder        custumers.Holder
	AgencyNumber  int
	AccountNumber int
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

func (c *CheckinAccount) Transfer(transferAmount float64, destinationAccount *CheckinAccount) string {
	if transferAmount <= c.balance && transferAmount > 0 {
		c.balance -= transferAmount
		destinationAccount.Deposit(transferAmount)
		return "Tranferência feita com sucesso"
	} else {
		return "Tranferência não concluída. Por favor verifique os valores"
	}
}

func (c *CheckinAccount) GetBalance() float64 {
	return c.balance
}
