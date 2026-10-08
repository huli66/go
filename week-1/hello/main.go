package main

import (
	"fmt"
	"math/rand"
	mrand "study/week-1/rand"
	"time"
)

const a = "hh"

func init() {
	const w = "world"
	fmt.Println("init", time.Now())
}

func main() {
	fmt.Println("hello", rand.Int())
	mrand.Rand()

	var (
		a bool   = true
		b string = "123"
		c int    = 12
	)

	d, a := false, false

	fmt.Println(a, b, c, d)
}
