package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main() {
	for i:=0;i<10;i++{
		wg.Add(1)
		go f(i)
	}
	wg.Wait()
}

func f(i int) {
	defer wg.Done()
	time.Sleep(time.Second* time.Duration(rand.Intn(3)))
	fmt.Println(i)
}