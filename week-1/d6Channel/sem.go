package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

type Result struct {
	Err  error
	Name string
	Size int
}

var running atomic.Int64

func download(name string) Result {
	fmt.Printf("%s 开始，（进行中:%d）\n", name, running.Add(1))
	defer running.Add(-1)
	t := rand.IntN(400) + 100
	time.Sleep(time.Duration(t) * time.Millisecond)

	if rand.IntN(5) == 0 {
		return Result{
			Name: name,
			Size: t,
			Err:  errors.New("download error"),
		}
	}
	return Result{
		Name: name,
		Size: t,
	}
}

func main() {
	files := []string{"11", "12", "13", "14", "15", "16", "17", "18", "19"}
	sem := make(chan int, 3)
	ch := make(chan Result)
	startTime := time.Now()

	var wg sync.WaitGroup

	for i, f := range files {
		wg.Go(func() {
			sem <- i
			defer func() { <-sem }()
			fmt.Println("start download\n", i)
			ch <- download(f)
		})
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	ok, fail, total := 0, 0, 0

	for r := range ch {
		if r.Err != nil {
			fail++
			fmt.Printf("%s 失败\n", r.Name)
		}
		ok++
		total++
		fmt.Printf("%s 成功， size  %d\n", r.Name, r.Size)
	}

	endTime := time.Since(startTime)

	fmt.Printf("total: %d, ok: %d, fail: %d, time: %d \n", total, ok, fail, endTime)
}
