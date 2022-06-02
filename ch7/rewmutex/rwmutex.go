package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var rwlock sync.RWMutex
var lock sync.RWMutex

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	rwlock.RUnlock()
	time.Sleep(time.Millisecond)
}
func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	rwlock.Unlock()
	time.Sleep(time.Millisecond * 5)

}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
