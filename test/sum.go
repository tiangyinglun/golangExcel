package main

import (
	"fmt"
	"runtime"
)

type person struct {
	Name  string
	Age   int
	Sex   int
	Addr  string
	Phone string
}

type Human struct {
	Name   string
	Age    int
	Weight int
}

type Skills []string

type Stu struct {
	Human
	Skills
	int
	Spc string
}

type Student struct {
	person
	Spc string
}

type Rect struct {
	Width, Height float64
}

func (r Rect) area() float64 {
	return r.Width * r.Height
}

type men interface {
	SayHi()
}

func (h Human) SayHi() {
	fmt.Println(h.Name)

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("+++++++++++")
		for i := 0; i < 10; i++ {
			fmt.Println("================")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	fmt.Println("*************")
	x, y := 1, 1
	for {
		fmt.Println("##############")
		select {
		case c <- x:
			fmt.Println("-----------")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fbTest() {
	c := make(chan int, 20)
	go fb(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fb(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func chanTest() {
	a := []int{7, 3, 2, 1, 7, 9, 5, 6}
	c := make(chan int)
	go Sum(a[:len(a)/2], c) //13
	go Sum(a[len(a)/2:], c) //27
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) //40
}

//求和
func Sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

//goroutine 并发

func goroutTest() {
	go Say("world")
	Say("hello")
}

func Say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func interfacetest() {
	Mark := Human{"hello", 12, 18}
	var i men = Mark
	i.SayHi()
}

func RectTest() {
	r1 := Rect{12.3, 12.5}
	fmt.Println(r1.area())
}
func structTest() {
	var team person
	team.Age = 2
	team.Addr = "a"
	team.Name = "hello"
	team.Phone = "1352245631"
	team.Sex = 1
	//	sam := person{Age: 12, Name: "aaa", Addr: "akdfljal", Phone: "13245678", Sex: 1}
	kak := person{"aaa", 15, 21, "akdfljal", "13245678"}

	ta, aa := Older(team, kak)
	fmt.Println(ta)
	fmt.Println(aa)
}

func Older(p1, p2 person) (person, int) {
	if p1.Age > p2.Age {
		return p1, p1.Age - p2.Age
	}
	return p2, p2.Age - p1.Age

}

//struct赋值
func structName() {
	var p person
	p.Name = "田朵朵"
	p.Age = 1
	p.Sex = 2
	p.Addr = "北京昌平"
	p.Phone = "1352213456"
	fmt.Println(p.Phone)
}

//指针类型
func Point1(a *int) int {
	*a = *a + 1
	return *a
}

//point

func Point(a int) int {
	a = a + 1
	return a
}

//Args(1, 2, 3, 4, 5)
//Args(s...)  //s := []int{1, 3}
func Args(arg ...int) { //变参
	for _, n := range arg {
		fmt.Println(n)
	}
}

func SunAndProduct(a, b int) (add, product int) {
	add = a + b
	product = a * b
	return
}

func mainTest() {
	x := 8
	y := 4
	z := 5

	max_xy := Max(x, y)
	max_xz := Max(x, z)
	fmt.Printf("max(%d,%d)=%d\n", x, y, max_xy)
	fmt.Printf("max(%d,%d)=%d\n", x, z, max_xz)
	fmt.Printf("max(%d,%d)=%d\n", y, z, Max(y, z))

}

func Max(a, b int) (c int) {
	if a > b {
		return a
	} else {
		return b
	}
}

func ManyArgs(a, b, c int) (m, t, k int) {
	return a, b, c
}

func ForMap() {
	data := make(map[string]string)
	data["one"] = "one"
	data["two"] = "two"
	data["three"] = "three"
	for k, v := range data {
		fmt.Println(k + "====" + v)
	}
}

func ForTest() {
	sum := 0
	for x := 0; x <= 10000; x++ {
		sum += x
	}
	fmt.Println(sum)

	sumt := 1
	for sumt < 100 {
		sumt += sumt
	}
	fmt.Println(sumt)
}

func TestIf() {
	a := 10
	if a < 10 {
		fmt.Println("i am big")
	} else if a > 10 {
		fmt.Println("i am small")
	} else {
		fmt.Println("i am this")
	}
}

func MapString() {
	m := make(map[string]string)
	m["one"] = "hello"
	m["two"] = "world"
	m1 := m
	m1["one"] = "the world you must stand up"
	fmt.Println(m1)

}

func MapTest() {
	//声明map
	//var num map[string]int
	num := make(map[string]int)
	num["one"] = 1
	num["two"] = 2
	//map值覆盖
	num["two"] = 12
	delete(num, "one") //删除map
	fmt.Println(num)
}

func Array() {
	var arr [5]int
	arr[0] = 1
	arr[2] = 3
	fmt.Printf("%d", arr[0])
	fmt.Println("\r\n")
	fmt.Printf("%d", arr[4])
	fmt.Println(arr)
}

func StringSum() (str string) {
	s := "12345"
	t := "6789"
	return s + t
}
func Cstring() (S2 string) {
	s := "hello"
	c := []byte(s) //  将字符串 s  转换为 []byte  类型
	c[0] = 'c'
	S2 = string(c) //  再转换回 string  类型
	return S2
}

func Test() (c int) {
	var a int
	a = 12
	return a
}

func Sums(a, b int) (c int) {
	return a + b
}
