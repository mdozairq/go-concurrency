package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Time: %v\n", time.Since(start))
	}()

	names := []string{"foo", "bar", "baz", "frr"}

	for _, name := range names {
		go callName(name)
    }

	time.Sleep(time.Second *2)

}

func callName(name string){
	fmt.Println("Hey "+name)
	time.Sleep(time.Second*2)
}