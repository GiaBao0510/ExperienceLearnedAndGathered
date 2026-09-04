package main

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(ammount int) {
	w.balance += ammount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(ammount int) error {
	if w.balance < ammount {
		return fmt.Errorf("Balance is not enough to debit")
	}

	w.balance -= ammount
	fmt.Println("Wallet balance debited successfully")
	return nil
}
