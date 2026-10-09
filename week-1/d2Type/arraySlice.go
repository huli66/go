package main

import (
	"fmt"
	"unsafe"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d ptr=%p %v \n", len(s), cap(s), unsafe.SliceData(s), s)
}

func arraySlice() {
	var arr [10]int // [n]T 是由 n 个 T 类型组成的数组，数组长度也是类型的一部分，所以数组不能调整大小
	arr[0] = 1

	fmt.Println(arr[0], arr[1], arr, arr[1:4]) // [1, 4)

	s1 := []int{1, 2, 3, 4, 5, 6}
	printSlice(s1)
	s2 := s1[:3]
	printSlice(s2)
	s2 = s2[:0]
	printSlice(s2)
	s2 = s2[:]
	printSlice(s2)
	s2 = s2[:5]
	printSlice(s2)
	// s2 = s2[:8] // 再切片不能超过长度，否则会报错
	s2 = s2[2:6]
	s3 := s2[:]
	printSlice(s2)
	s2 = append(s2, 2) // 超出容量自动扩容
	printSlice(s2)     // 扩容之后指针指向新的地址
	printSlice(s3)

	fmt.Println(*(unsafe.SliceData(s3)))
	fmt.Println(*(unsafe.SliceData(s2)))

	var s4 []int
	if s4 == nil { // 编译器提示恒为真
		fmt.Println("nil")
	}
	printSlice(s4)

	s5 := make([]int, 10)
	printSlice(s5)
	s6 := make([]int, 0, 10)
	printSlice(s6)

	// s7 := []int{1, 2, 0, 4, 5, 0, 0, 8, 0, 0}

	for i, v := range s5 {
		fmt.Println("range", i, v)
	}
	a8 := [7]int{1, 1, 1, 1, 1, 1, 1}
	s8 := a8[:]
	for i, _ := range s8 {
		if s8[i] > 0 && i < len(s8)-1 {
			s8[i+1] = -1
		}
	}
	fmt.Println(s8, a8)
}

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i, _ := range p {
		p2 := make([]uint8, dx)
		p[i] = p2

		for i2, _ := range p2 {
			p2[i2] = uint8(i ^ i2)
			// p2[i2] = uint8(i * i2)
			// p2[i2] = uint8((i + i2) / 2)
		}
	}

	return p
}
