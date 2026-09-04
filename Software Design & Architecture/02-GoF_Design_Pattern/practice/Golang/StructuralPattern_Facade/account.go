package main

import "fmt"

type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{name: name}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name does not match")
	}

	fmt.Println("Account name matches")
	return nil
}