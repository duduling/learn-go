package main

import (
	"fmt"

	"github.com/duduling/learngo/mydict"
)

func main() {
	// #region - Account Tutorial
	// account := accounts.NewAccount("dudu")
	// account.Deposit(10)
	// err := account.Withdraw(10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account)
	// #endregion

	//  #region
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	//  #endregion
}
