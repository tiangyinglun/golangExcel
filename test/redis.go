package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//redis 连接池

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "10.111.23.230:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// 生成连接池
var pool = newPool()

func main() {
	Lpush()
}

func Lpush() {
	c := pool.Get()
	defer c.Close()
	c.Send("LPUSH", "albums", "hello")
	c.Send("LPUSH", "albums", "world")
	c.Send("LPUSH", "albums", "i want fly")
}

func Exists() {
	c := pool.Get()
	defer c.Close()
	exists, _ := redis.Bool(c.Do("EXISTS", "a"))
	fmt.Printf("%#v\n", exists)
}

func PoolSet() {
	c := pool.Get()
	defer c.Close()
	ok, err := c.Do("SET", "dadasasd", "world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)
}

func Set() {
	c, err := redis.Dial("tcp", "10.111.23.230:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	v, err := c.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println(err)
		return
	}

	v, err = redis.String(c.Do("GET", "a"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	defer c.Close()
}
