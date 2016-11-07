package main

import (
	"fmt"
	//	"time"
)

func main() {
	i := 0
	m := 0
	for i < 10000000000 {
		i++
		m++
	}
	fmt.Println(m)
}
