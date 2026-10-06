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
}
