package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

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

	// #region - My dict
	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// dictionary.Search(baseWord)
	// dictionary.Delete(baseWord)
	// word, err := dictionary.Search(baseWord)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(word)
	// }
	// #endregion

	// #region - URL
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
	// #endregion
}

func hitURL(url string) error {
	fmt.Println("Checking", url)

	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
