package main

import "fmt"

func printByte1(str string, ch chan byte, c1 chan bool, c2 chan bool) {
	for {
		<-c1
		fmt.Printf("str is %v is %v\n", str, string(<-ch))
		c2 <- true
	}
}

func printByte2(str string, ch chan byte, c1 chan bool, c2 chan bool) {
	for {
		<-c2
		fmt.Printf("str is %v is %v\n", str, string(<-ch))
		c1 <- true
	}
}

func main() {
	ch := make(chan byte)
	c1, c2 := make(chan bool), make(chan bool)
	strs := "123456"

	go printByte1("thread1", ch, c1, c2)
	go printByte2("thread2", ch, c1, c2)

	c1 <- true
	for i := 0; i < len(strs); i++ {
		ch <- strs[i]
	}
	close(ch)
	close(c1)
	close(c2)
}
