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
}
