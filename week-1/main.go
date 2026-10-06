package main

import (
	"fmt"
	"study/week-1/rand"
)

func init() {
	fmt.Println("main week1 init")
}

func main() {
	fmt.Println("main week1")
	rand.Rand()

	var a = rand.NewUser("huli", 12)
	fmt.Println(a.Name, a.Show()) // a.age 会报错 undefined，因为 age 属性是小写的，不会暴露

	const k = 1
	var i int = k
	fmt.Println("int", i)
	var f float64 = k
	fmt.Println("float64", f)
	var b byte = k
	fmt.Println("byte", b)
	var ( // 连续声明多个不同类型变量，避免写很多 var
		v1 bool
		v2 string
		v3 int
		v4 float64
	)
	fmt.Println(v1, v2, v3, v4)
	fmt.Println("hhh")
}
