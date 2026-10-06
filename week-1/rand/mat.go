package rand

import (
	"fmt"
)

func init() {
	fmt.Println("mat init")
}

type user struct { // 未导出，无法直接使用
	Name string // 导出，能取值，gopls 有提示
	age  int    // 未导出，不能取值，会 undefined，gopls 没有提示
}

func NewUser(n string, a int) user {
	return user{n, a} // 通过构造函数使用 user 类型
}

func (u user) Show() string {
	return u.Name
}

func Rand() {
	fmt.Println("rand random")
}

func Hello(name string, age int) string {
	return name
}

// 类型可以连续，可以返回多个值（和 TS 不一样，不需要包裹一层）
func Hi(x, y string) (string, string) { return x, y }

// 命名返回值，会被视为在函数顶部声明了变量
func Ha(a, b int) (sum, y int) {
	sum = a + b
	y = a * b
	return // 裸返回，不带参数的 return， 会返回命名返回值
	// 只建议在短函数里使用裸返回，长函数里使用会影响可读性
}

func twosum(x, y int) int {
	x, b := 1, 2
	var sum int         // var 声明变量，可以在包级别也可以在函数级别，必须带类型
	fmt.Println(sum, b) // 0
	return sum
}
