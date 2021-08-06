package main

import (
	"fmt"
	"time"
)

func main() {
	// go func() {
	// 	time.Sleep(3 * time.Second)

	// }()
	for true {
		fmt.Println("begin")
		con := make(chan bool)
		// t := rand.Intn(150)
		timer1 := time.NewTimer((100 + 300) * time.Millisecond)
		go func() {
			timer1.Stop()
			//con <- true
		}()

		//go func() {
		//	for ;true; {
		//		fmt.Println("avoid sleep")
		//	}
		//}()
		select {
			case <-timer1.C:
				fmt.Println("end")
			case <-con:
				fmt.Println("con")

		}
	}
}
