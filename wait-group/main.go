package main

import (
	"fmt"
	"sync"
	// "time"
)


func main() {
	var caller sync.WaitGroup
	fmt.Println("Hey Main Thread 1")
	names := []string{"ABC", "XYZ", "PQR", "LMN"}
	caller.Add(len(names))
	for _, name := range names {
		go callName(name, &caller)
	}
	caller.Wait()
	fmt.Println("Hey Main Thread 2")
}

func callName(name string, callback *sync.WaitGroup){
	// time.Sleep(time.Second*2)
	fmt.Println("Hey", name)
	callback.Done()
}