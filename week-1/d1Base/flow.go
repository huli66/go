package main

import (
	"fmt"
	"math"
	"runtime"
)

func completeLoop(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("complete loop", i)
	}
}

func whileLoop(x int) {
	n := 0
	for n < x {
		n = n + 1
		fmt.Println("while ", n)
	}
}

func rangeLoop(x int) {
	for ; x < 100; x++ {
	}
}

func mySqrt(x int) float64 {
	if x < 0 {
		return 0
	}
	return 0
}

func myIf(x int) int {
	if v := x * 2; v > 1 {
		return v //因为 if 前置语句中声明的变量作用域只持续到 if 语句结束，包括 else
	} else if m := 3; m > v {
		return m * v //
	}
	// return v; // undefined v
	return 0
}

/**
 * 牛顿法求平方根
 */
func newDun(x float64) float64 {
	z := float64(1)
	r := math.Sqrt(x)
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z) // 牛顿法，求任意数字的开方，暂时不知道怎么证明
		fmt.Println("newDun", i, z)
		if z-r < 0.0000001 {
			return z // 是否会中断循环
		}
	}
	return 0
}

func mySwitch() {
	defer newDun(newDun(4)) // 参数会立即执行，但是外层会延迟
	defer newDun(9)
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "drain":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.\n", os)
	}

}

func namedReturn() int {
	n := 4
	defer func() { n *= 2 }()
	return n // 直接 return 常量会
}

func main() {
	// completeLoop(5)
	// whileLoop(3)

	// fmt.Println("1", newDun(1))
	// fmt.Println("2", newDun(2))
	// fmt.Println("3", newDun(3))
	// fmt.Println("4", newDun(4))
	// fmt.Println("5", newDun(5))
	// fmt.Println("6", newDun(9))

	// mySwitch()

	// fmt.Println(namedReturn())

	for i, v := range "hello" {
		fmt.Println(i, v)
	}
}
