package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Flag bool   `json:"flag"`
}

func (s Student) Greet() {
	f := "Sir"
	if s.Flag {
		f = "Madon"
	}
	fmt.Println("hello ", f, s.Name)
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
	fmt.Println("newName", s)
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	fmt.Println("c", c.R)
	return 3.14 * c.R * c.R
}

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	type plain Person // 定义一个新类型，不是别名
	// 新类型只继承字段不继承方法
	return fmt.Sprintf("%#v, %+v, %v (%v years)\n", p, plain(p), p.Name, p.age)
}

func Index[T comparable](a []T, x T) int {
	for i, v := range a {
		if x == v {
			return i
		}
	}
	return -1
}

func methodTest() {
	fmt.Println("method")
	s := Student{
		12,
		"huli",
		false,
	}
	s.Greet()
	s.ChangeName("Li")
	s.Greet()
	Student.Greet(s)

	j, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(j))

	var c Circle
	c.Area()
	fmt.Println("c", c)

	type Shape interface {
		Area() float64
	}
	var d Circle = Circle{3}
	var e Shape = &d
	e.Area() // e 只能读取方法，不会暴露里面的值
	// 方法集
	fmt.Println(e)

	type Zeroi interface{}
	var zzz Zeroi = 0
	fmt.Println(zzz)

	// 类型断言
	var h any = "hello"
	if t, ok := h.(float64); ok {
		fmt.Println("打印", t)
	} else {
		fmt.Println("zzz 没有保存 Zeroi 类型，则 t 为该类型零值", t)
	}

	p1 := Person{"huli", 18}
	p2 := Person{"li", 19}
	fmt.Println(p1, p2)

	ss1 := []string{"hell0", "hi", "yeah"}
	idx := Index(ss1, "hi")
	fmt.Println(idx)
}
