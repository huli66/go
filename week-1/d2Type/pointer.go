package main

import "fmt"

// 指针和 struct

func pointer() {
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

func pointer2() {
	a, b := 1, 2
	c, d := &a, &b
	var e *int

	fmt.Println(c, d, e)
	fmt.Println(*c, *d) // 空指针不能解引用
	fmt.Printf("%p %p %p\n", c, d, e)

	type MyVec struct {
		X int
		Y *int
	}

	v := MyVec{a, c}
	fmt.Println(v.X, v.Y, *v.Y)
	fmt.Println("部分声明，不要换行", MyVec{X: 1})
}
