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
	Succeeded() bool
	SetSucceeded(value bool)
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

func (bac *BankAccountCommand) Succeeded() bool {
	return bac.succeeded
}

func (bac *BankAccountCommand) SetSucceeded(value bool) {
	bac.succeeded = value
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (bac *CompositeBankAccountCommand) Call() {
	for _, cmd := range bac.commands {
		cmd.Call()
	}
}

func (bac *CompositeBankAccountCommand) Undo() {
	for idx := range bac.commands {
		cmd := bac.commands[len(bac.commands)-idx-1]
		cmd.Undo()
	}
}

func (bac *CompositeBankAccountCommand) Succeeded() bool {
	succeeded := true
	for _, cmd := range bac.commands {
		if !cmd.Succeeded() {
			succeeded = false
			break
		}
	}
	return succeeded
}

func (bac *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range bac.commands {
		cmd.SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands,
		NewBankAccountCommand(from, Withdraw, amount),
		NewBankAccountCommand(to, Deposit, amount),
	)
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
			continue
		}

		cmd.SetSucceeded(false)
	}
}

func main() {
	from := BankAccount{100}
	to := BankAccount{0}
	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from.balance, to.balance)

	mtc.Undo()
	fmt.Println(from.balance, to.balance)
}
