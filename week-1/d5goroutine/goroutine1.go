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

func download(name string, ch chan<- int) {
	t := rand.IntN(4000) + 1000
	time.Sleep(time.Duration(t) * time.Millisecond)
	ch <- t // 写入
	fmt.Println("complete", t, name)
}

func gen(n int) <-chan int {
	ch := make(chan int)

	go func() {
		// 提前登记 defer，就算 panic 也会执行
		defer close(ch)
		for i := range n {
			fmt.Println("write", i)
			ch <- i
		}
	}()
	return ch
}

func main1() {
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
		// go download(f, ch)
		wg.Go(func() { download(f, ch) })
	}
	go func() {
		fmt.Println("close1")
		wg.Wait()
		fmt.Println("close2")
		close(ch)
	}()
	fmt.Println("close3")
	// close(ch)
	endTime := time.Since(start)
	total := 0
	for v := range ch {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("v", v)
	}
	// ch <- 2
	// for range files {
	// 	time.Sleep(2000)
	// 	total += <-ch
	// }
	// ch <- 2
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("start", start, "end", endTime, total)

}

func dl(name string)

func main() {
	ch2 := gen(5)
	for chc := range ch2 {
		fmt.Println("read", chc)
	}
}
