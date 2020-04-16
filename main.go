package main

import (
	"fmt"

	"github.com/duduling/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("dudu")
	fmt.Println(account)
}
