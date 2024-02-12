package main

import (
	"log"
	"patterns/code"
)

// Смысл в том, что мы фасадом скрываем внутрянку (все процессы) от пользователя

func main() {
	walletFacade := code.NewWalletFacade("abc", 1234)
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
