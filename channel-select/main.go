package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	defer func() {
		fmt.Println("Process complete after:", time.Since(start))
	}()
	candidate1, candidate2 := make(chan string), make(chan string)
	
	go electMonitor(candidate1, "ABC")
	go electMonitor(candidate2, "XYZ")
	// time.Sleep(time.Second * 2)
	select {
	case message := <-candidate1:
		fmt.Println(message,"is Winner!")
	case message := <-candidate2:
		fmt.Println(message,"is Winner!")
	default:
		fmt.Println("None Elected!")
	}
}

func electMonitor(candidate chan string, name string){
	time.Sleep(time.Second * 2)
	candidate <- name
}