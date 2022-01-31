package main

import (
	"fmt"
	"runtime"
	"sync"
)

var cont int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			v := cont
			runtime.Gosched()
			v++
			cont = v
			mu.Unlock()
			fmt.Println("cont =", cont)
			wg.Done()
		}()
	}
	wg.Wait()
}
