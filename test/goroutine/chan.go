package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	aa := make([]int, 20000)
	go func() {
		for i := 0; i < 20000; i++ {
			aa[i] = <-c
		}
	}()
	num(10000, c)

	//num2(10000, c)

	fmt.Println(aa)

}

func num(s int, c chan int) {
	x := 0
	for x < s {
		c <- x
		x++
	}
}

// func num2(s int, c chan int) {
// 	x := s
// 	for x > 0 {
// 		c <- x
// 		x--
// 	}
// }
