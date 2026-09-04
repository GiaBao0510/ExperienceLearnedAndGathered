package main

import "fmt"

type WalletFacade struct {
	account       *Account
	wallet        *Wallet
	securityCode  *SecurityCode
	nottification *Notification
	ledger        *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create wallet facade")
	walletFacade := &WalletFacade{
		account:       newAccount(accountID),
		securityCode:  newSecurityCode(code),
		wallet:        newWallet(),
		nottification: &Notification{},
		ledger:        &Ledger{},
	}

	fmt.Println("Wallet facade created")
	return walletFacade
}

// Thêm tiền vào ví
func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, ammount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	// thực hiện thêm tiền
	w.wallet.creditBalance(ammount)
	w.nottification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", ammount)
	return nil
}

// Lấy tiền ra khỏi ví
func (w *WalletFacade) debitMoneyFromWallet(accountID string, securityCode int, ammount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	// thực hiện lấy tiền ra khỏi ví
	err = w.wallet.debitBalance(ammount)
	if err != nil {
		return err
	}
	
	w.nottification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "debit", ammount)
	return nil
}