package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (ba *BankAccount) Withdraw(amount int) bool {
	if ba.balance-amount >= overdraftLimit {
		ba.balance -= amount
		fmt.Println("Withdrew ", amount, "\b, balance is now ", ba.balance)
		return true
	}
	return false
}

func (ba *BankAccount) Deposit(amount int) {
	ba.balance += amount
	fmt.Println("Deposited ", amount, "\b, balance is now ", ba.balance)
}

type Command interface {
	Call()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (bac *BankAccountCommand) Call() {
	switch bac.action {
	case Deposit:
		bac.account.Deposit(bac.amount)
		bac.succeeded = true
	case Withdraw:
		bac.succeeded = bac.account.Withdraw(bac.amount)
	}
}

func (bac *BankAccountCommand) Undo() {
	if !bac.succeeded {
		return
	}

	switch bac.action {
	case Deposit:
		bac.account.Withdraw(bac.amount)
	case Withdraw:
		bac.account.Deposit(bac.amount)
	}	
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba.balance)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	cmd2.Call()
	fmt.Println(ba.balance)
	cmd2.Undo()
	fmt.Println(ba.balance)
}
