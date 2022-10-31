package main

import "fmt"

func userCh() int {

	usCh := make(chan string)
	go func(){
	usCh <- "chemistry"
	}()

		user := <-usCh

	fmt.Println(user)
	return 0

}
