package main

import (
	"fmt"
	"time"
)

//channel without goroutine will panic a error
 
func main(){
	ch := make(chan string)
	fmt.Println(time.Now(), "taking a nap!")

	go func(){
		time.Sleep(time.Second *5)

		ch <- "Hello from Channel"
	}()

	v := <-ch
	fmt.Println("Waiting for message...\n", v)
	fmt.Println(time.Now(), "completed task!")
}