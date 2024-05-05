package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Process complete after:", time.Since(start))
	}()
	time.Sleep(time.Second * 2)
	channel := make(chan bool)
	name := "Ozair"
	go xyz(name, channel)
	channel <- false

	fmt.Println("Hello Value: ", <-channel)
}

func xyz(name string, channel chan bool) {
	fmt.Println("Hey ", name)
	time.Sleep(time.Second * 2)
	channel <- true
}
