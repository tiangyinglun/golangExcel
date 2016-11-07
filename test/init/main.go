package main

import (
	"fmt"
	"init/demo"
)

func main() {
	myConfig := new(demo.Config)
	myConfig.InitConfig("./config.ini")
	fmt.Println(myConfig.Read("helloworld", "addr"))
	fmt.Println(myConfig.Read("helloworld", "port"))
	//fmt.Printf("%v", myConfig.Mymap)
}
