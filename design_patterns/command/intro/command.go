package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (ba *BankAccount) Withdraw(amount int) {
	if ba.balance - amount >= overdraftLimit {
		ba.balance -= amount
		fmt.Println("Withdrew ", amount, "\b, balance is now ", ba.balance)
	}
}

func (ba *BankAccount) Deposit(amount int) {
	ba.balance += amount
	fmt.Println("Deposited ", amount, "\b, balance is now ", ba.balance)
}

type Command interface {
	Call()
}

type Action int
const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action Action
	amount int
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount}
}

func (bac *BankAccountCommand) Call() {
	switch bac.action {
	case Deposit:
		bac.account.Deposit(bac.amount)
	case Withdraw:
		bac.account.Withdraw(bac.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba.balance)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba.balance)
}
