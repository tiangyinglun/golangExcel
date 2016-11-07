package main

import (
	"fmt"
)

func main() {
	fb2()
}

func fb2() {
	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("-----------")
			fmt.Println(<-c4)
		}
		quit <- 0
	}()

	fibonacci2(c4, quit)
}

//第一次到select 阻塞
//第二取
//取不到 select赋值 立马取出来
//然后走select case里的
//给了 c 能立马别取走 感觉像走后门

//++++
//----
//0
//-----
//*****
//+++
//***
//+++
//1
//----
//1
//-----
//****
//+++
//****
//+++
//2
//-----
//3
//-----
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("++++++")
		select {
		case c <- x:
			fmt.Println("*****************")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fb1() {
	c := make(chan int, 1)
	go fibonacci(10, c)
	for i := range c {
		fmt.Println("----------")
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println("+++")
		c <- x
		x, y = y, x+y
	}

	close(c)
}
