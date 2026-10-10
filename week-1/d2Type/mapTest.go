package main

import (
	"fmt"
	"maps"
)

func bibaoAdd() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fib() func() int {
	f := 1
	s := 1
	return func() int {
		s = s + f
		f = s - f
		return s
	}
}

func mapTest() {
	var m map[string]int
	if m == nil {
		fmt.Println("nil")
	}
	fmt.Println(m)

	m1 := make(map[string]int)
	fmt.Println(m1)
	if m1 == nil {
		fmt.Println("make nil")
	}
	fmt.Printf("%#v %#v\n", m, m1)
	m1["sss"] = 1
	fmt.Println(m1)

	// 可以初始化的时候带上字面量
	var m2 = map[string]int{} // 效果和 make 创建一样
	m3 := make(map[string]int)
	m4 := make(map[string]int, 10)
	fmt.Printf("%#v %#v %#v\n", m2, m3, m4)
	fmt.Printf("%p %p %p\n", m2, m3, m4)
	if maps.Equal(m2, m3) { // 内容比较
		fmt.Println("equal")
	}
	if maps.Equal(m2, m4) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	m5 := make(map[string]any)
	m5["key"] = 2
	val, ok := m5["key"]
	// 如果 key 这个键存在，ok 为 true，否则 false
	//  键不存在，则 val 的值是该元素类型的零值
	fmt.Println(val, ok, m5)

	switch ty := val.(type) {
	case int:
		fmt.Println("int", ty)
	default:
		fmt.Println("any", ty)
	}

	b1, b2 := bibaoAdd(), bibaoAdd()

	for i := range 10 {
		fmt.Println(b1(i), b2(i*2))
	}
	fn := fib()
	for i := range 5 {
		fmt.Println(i, fn())
	}
}
