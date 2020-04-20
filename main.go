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
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err = dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	//  #endregion
}
