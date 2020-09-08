package main

import (
	"fmt"

	"github.com/tonnac/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ewew")
	fmt.Println(account.Balance())
	account.Deposit(32)
	fmt.Println(account.Balance())
	err := account.Withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
