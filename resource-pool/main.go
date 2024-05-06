package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryPeices int
	memPool := &sync.Pool{
		New: func() interface{} {
			memoryPeices++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const workerSize = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerSize)

	for i := 0; i < workerSize; i++ {
		go func() {
			// memPool.Get()
			mem := memPool.Get().(*[]byte)
			fmt.Sprintln("time to allocate resources")
			memPool.Put(mem)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Number of memory pieces created are:", memoryPeices)
}
