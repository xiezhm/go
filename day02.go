package main

import (
	"fmt"
)

//
//func main()  {
//	for i:=0;i<=1000;i++{
//		println(i,"我是go语言哦!")
//	}
//}

func main()  {
	var a float64
	for i:=528;i<=535;i++{
		a = 64986/float64(i)
		fmt.Print(a,"\n")
	}
	fmt.Print("\n")
	for m:=121;m<=123;m++{
		a = 64986/float64(m)
		fmt.Print(a,"\n")
	}
}
