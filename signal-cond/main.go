package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// type Cond struct{
// 	noCopy noCopy
// }

// type WaitJar interface{

// }

func main(){
	gettingReadyForMission()
}

var ready bool


func gettingReadyForMissionWithCond(){
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	// go gettingReady()
	workIntervals := 0
	cond.L.Lock()

	for !ready {
		// time.Sleep(time.Second * 5)
		workIntervals++
		cond.Wait()
	}

	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals \n", workIntervals)
}

func gettingReadyForMission(){
	go gettingReady()
	workIntervals := 0
	for !ready {
		time.Sleep(time.Second * 5)
		workIntervals++
	}

	fmt.Printf("We are now ready! After %d work intervals \n", workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond){
	sleep()
	ready=true
	cond.Signal()
}

func gettingReady(){
	sleep()
	ready=true
}

func sleep(){
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}