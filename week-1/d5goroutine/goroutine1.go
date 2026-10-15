package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func say(s string) {
	for i := range 3 {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i, time.Now().Format("15:04:05.000000"), s)
	}
}

func download(name string, ch chan int) {
	t := rand.IntN(4000) + 1000
	time.Sleep(time.Duration(t) * time.Millisecond)
	ch <- t // 写入
	fmt.Println("complete", t, name)
}

func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		fmt.Println("wg", i)
		wg.Go(func() {
			say("hello")
		})
	}
	wg.Wait() // 阻塞当前 goroutine
	fmt.Println("complete start")

	start := time.Now()
	files := []string{"file1", "file2", "file3"}
	ch := make(chan int)
	for _, f := range files {
		fmt.Println("download", f)
		go download(f, ch)
		// wg.Go(func() { download(f, ch) })
	}
	wg.Wait()
	endTime := time.Since(start)
	total := 0
	// ch <- 2
	for range files {
		time.Sleep(2000)
		total += <-ch
	}
	ch <- 2
	fmt.Println("start", start, "end", endTime, total)
}
