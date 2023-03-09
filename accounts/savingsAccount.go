package accounts

import (
	"bank/custumers"
)

type SavingsAccount struct {
	Holder                                 custumers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (s *SavingsAccount) ToWithdraw(withdrawalAmount float64) string {
	canWithdraw := withdrawalAmount >= 0 && withdrawalAmount <= s.balance
	if canWithdraw {
		withdrawalAmount -= s.balance
		return "Saque feito com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (s *SavingsAccount) Deposit(amountDeposited float64) (string, float64) {
	if amountDeposited > 0 {
		s.balance += amountDeposited
		return "Depósito feito com sucesso", s.balance
	} else {
		return "Depósito não realizado. Verifique os valores", s.balance
	}
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}
