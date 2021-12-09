package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func Withdraw(ba *BankAccount, amount int) bool {
	if ba.balance-amount >= overdraftLimit {
		ba.balance -= amount
		fmt.Println("Withdrew ", amount, "\b, balance is now ", ba.balance)
		return true
	}
	return false
}

func Deposit(ba *BankAccount, amount int) {
	ba.balance += amount
	fmt.Println("Deposited ", amount, "\b, balance is now ", ba.balance)
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands,
		func() {
			Deposit(ba, 100)
		},
	)
	commands = append(commands,
		func() {
			Withdraw(ba, 25)
		},
	)

	for _, cmd := range commands {
		cmd()
	}

	fmt.Println(ba)
}
