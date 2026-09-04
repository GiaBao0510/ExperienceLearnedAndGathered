package main

import "fmt"

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("user01",6868)

	err := walletFacade.addMoneyToWallet("user01",6868,100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n--------------------")
	err = walletFacade.debitMoneyFromWallet("user01",6868,100)
	if err != nil {
		fmt.Println(err)
	}
}