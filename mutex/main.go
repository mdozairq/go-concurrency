package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	rwLock sync.RWMutex
	count int
)

func main(){
	readAndWrite()
}

func readAndWrite(){
	
	go read()
	go read()
	go read()
	go read()
	go write()
	go read()
	go read()
	

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read(){
 rwLock.Lock()
 defer rwLock.Unlock()
 fmt.Println("Read Locking...")
 time.Sleep(1 * time.Second)
 fmt.Println("Read Unlocking...")
}

func write(){
	rwLock.Lock()
	defer rwLock.Unlock()
	fmt.Println("Write Locking...")
	time.Sleep(1 * time.Second)
	fmt.Println("Write Unlocking...")
}

func basic(){
	maxRange := 1000
	for i:=0; i<maxRange; i++{
		go increment()
	}

	time.Sleep(time.Second*1)
	fmt.Println("total goroutines: ", count)
}

func increment(){
	lock.Lock()
	count++
	lock.Unlock()
}