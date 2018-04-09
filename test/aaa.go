package main

import "fmt"

func Rec(i int)  {
	if i == 10 {
		return
	}
	Rec(i+1)
	fmt.Printf("%d",i)
}

func main() {
	Rec(0)
}