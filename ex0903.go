package main

import (
	"fmt"
	"runtime"
	"sync"
)

var cont int = 0
var wg sync.WaitGroup

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			v := cont
			runtime.Gosched()
			v++
			cont = v
			fmt.Println("cont =", cont)
			wg.Done()
		}()
	}
	wg.Wait()
}
