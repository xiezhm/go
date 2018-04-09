//package main
//
//import "fmt"
//
//func main() {
//	//for i :=1; i <= 10; i++  {
//	//	fmt.Printf("i de zhi:%d\n",i)
//	//
//	//}
//	for i :=1; i<=100; i++{
//		switch {
//		case i%5==0 && i%3 == 0:
//			fmt.Printf("FizzBuzz %d\n",i)
//		case i%3 == 0:
//			fmt.Printf("Fizz     %d\n",i)
//		case i%5 == 0:
//			fmt.Printf("Buzz     %d\n",i)
//		default:
//			fmt.Printf("         %d\n",i)
//		}
//	}
//}
//
//
//
//package main
//
//import "fmt"
//
//func main() {
//	const (
//		FIZZ =3
//		BUZZ =5
//	)
//	var p bool
//	for i := 1; i < 100; i++ {
//		p = false
//		if i%FIZZ == 0 {
//			fmt.Printf("Fizz")
//			p = true
//		}
//		if i % BUZZ ==0 {
//			fmt.Printf("Buzz")
//			p = true
//		}
//		if !p {
//			fmt.Printf("%v",i)
//		}
//		fmt.Println()
//	}
//
//}
package main

import "fmt"

func Rec(i int)  {
	if i == 10 {
		return
	}
	Rec(i+1)
	fmt.Printf("%d",i)
}

func main()  {
	Rec(0)
}






















