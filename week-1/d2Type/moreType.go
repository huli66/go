package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	var (
		a int    = 12
		b string = "123"
	)
	var c *int // 空值 0x0
	fmt.Printf("指针的值, %p\n", c)
	fmt.Println(a, b, c)

	c = &a
	fmt.Printf("%p\n", c)
	fmt.Println("c ", *c)
	a = 16
	fmt.Println(*c + 3)

	d := &c
	fmt.Println("指针的指针", *d)
	fmt.Printf("%p\n", d)

	type Vec struct {
		x int
		y *int
	}
	fmt.Println("结构体取值", *(Vec{1, c}).y, c)
	var v Vec = Vec{x: 1}
	pV := &v
	pV.x = 1e9
	fmt.Println("部分初始化", pV.y, "无需显示解引用", pV)
}
