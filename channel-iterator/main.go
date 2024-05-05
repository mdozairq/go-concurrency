package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Process complete after:", time.Since(start))
	}()
	// time.Sleep(time.Second * 2)
	channel := make(chan string)
	names := []string{"A", "B", "C"}
	studentRange := 3
	
	go testMark(names, channel, studentRange)
        
	c:=0
	for {
		v, _ := <-channel
		
		if c>2 {
			break
		}
		fmt.Println(v)
		c++
	}
}
	

func xyz(name string, channel chan bool) {
	fmt.Println("Hey ", name)
	time.Sleep(time.Second * 2)
	channel <- true
}

func testMark(names []string,channel chan string, studentRange int){
	rand.Seed(time.Now().UnixNano())
	for _, name := range names{
		score := rand.Intn(10)
		channel <- fmt.Sprint(name," scored: ",score," marks")
	}
}

