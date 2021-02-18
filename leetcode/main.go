package main

import "fmt"

func main(){
	t := "ABC"
	for _, v := range t{
		fmt.Println(byte(v))
	}
}