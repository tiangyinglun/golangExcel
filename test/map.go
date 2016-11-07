package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test3()
}

func test1() {
	var mp map[int]string

	mp = make(map[int]string)
	mp[1] = "hello world"
	mp[2] = "nihao shijie"
	fmt.Println(mp)
}

func test2() {
	mp := make(map[string]string)
	mp["one"] = "tianduoduo"
	mp["two"] = "duoduo"
	mp["three"] = "健康"
	mp["four"] = "快乐"
	fmt.Println(mp)

	for k, v := range mp {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func test3() {
	mp := make(map[string]interface{})
	mp["one"] = "hello"
	mp["two"] = 123
	fmt.Println(mp)
	js, err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
