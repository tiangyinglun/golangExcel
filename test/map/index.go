package main

import (
	"fmt"
)

type user struct {
	name string
}

type userInfo struct {
	name string
	age  int
	sex  int
}

type userTs struct {
	class string
	score int
}

type people struct {
	user
	userTs
	think string
}

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner string
	}
}

func main() {
	var p people
	p.name = "tianduoduo"
	p.score = 100
	p.class = "grade one"
	p.think = "happy"
	fmt.Println(p)
}

func fTest() {
	var f File
	f.name = "test.txt"
	f.size = 1024
	f.attr.owner = "tian yinglun"
	f.attr.perm = 777
	fmt.Println(f)
}

func structTest() {
	var a userInfo
	a.name = "小明"
	a.age = 12
	a.sex = 1
	fmt.Println(a)
}

func mapTest() {
	m := map[int]user{
		1: {"user1"},
	}

	fmt.Println(m[1])
	u := m[1]
	u.name = "hello world"
	m[1] = u
	fmt.Println(m)
}
