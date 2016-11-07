package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
)

func main() {
	crc322()
}

func bases64() {
	str := "每当我找不到存在的意义，每当我迷失在黑夜里"

	strb := []byte(str)
	aa := base64.StdEncoding.EncodeToString(strb)
	fmt.Println("---------------------------")
	fmt.Println(aa)
	data, err := base64.StdEncoding.DecodeString(aa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func jsons() {
	str := []string{"hello world", "laopo", "nv er", "money", "self"}

	jsonst, _ := json.Marshal(str)
	var strs []string
	err := json.Unmarshal(jsonst, &strs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strs)

}

func crc322() {
	str := "1234"
	nyt := []byte(str)
	table := crc32.MakeTable(0x04c11db7)
	hash := crc32.New(table)
}
