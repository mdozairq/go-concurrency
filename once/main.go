package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	isPassed bool
	once sync.Once
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			if getPass() {
				once.Do(setResult)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	checkResult()
	fmt.Println("Result Found!")
}

func checkResult() {
	if isPassed {
		fmt.Println("Student Passed Successfully")
	} else {
		fmt.Println("Student Failed!")
	}
}

func setResult() {
	isPassed = true
	fmt.Println("Setting Pass result")
}

func getPass() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
