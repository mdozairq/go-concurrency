package main

import (
	"fmt"
	"time"
)

//channel without goroutine will panic a error

func main() {
	ch := make(chan int, 2)
	exit := make(chan struct{})
	fmt.Println(time.Now(), "taking a nap!")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending...")
			ch <- i
			fmt.Println(time.Now(), i, "sent!")
			time.Sleep(time.Second * 2)
		}
		// time.Sleep(time.Second * 2)
		fmt.Println(time.Now(),  "All completed leaving!")

		close(ch)
	}()

	go func() {
		// for {
		// 	select{
		// 	case v, open := <-ch:
		// 		if !open {
		// 			close(ch)
		// 			return
		// 		}
		// 		fmt.Println(time.Now(), "recieved: ", v)
		// 	// default:
		// 	// 	fmt.Println(time.Now(), "Nothing Happening")
		// 	}
		// }
		// close(ch)

		for v := range ch {
			fmt.Println(time.Now(), "recieved: ", v)
		}

		close(exit)
	}()


	<-exit
	fmt.Println(time.Now(), "completed task!")
}
