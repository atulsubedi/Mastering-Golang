package main

import (
	"fmt"
	"time"
)

func userCh() int {

	usCh := make(chan string)
	go func(){
		time.Sleep(1*time.Second)
	usCh <- "chemistry"
	}()

		user := <-usCh

	fmt.Println(user)
	return 0

}
