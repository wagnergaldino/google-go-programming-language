package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var cont int64 = 0
var wg sync.WaitGroup

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&cont, 1)
			runtime.Gosched()
			fmt.Println("cont =", atomic.LoadInt64(&cont))
			wg.Done()
		}()
	}
	wg.Wait()
}
