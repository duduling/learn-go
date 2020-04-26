package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	c := make(chan bool)
	people := [2]string{"noci", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
