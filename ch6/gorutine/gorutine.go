package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var b chan int
	fmt.Println(b)
	b = make(chan int, 1)
	fmt.Println(b)
	// b = make(chan int, 16)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	x := <-b
	// 	fmt.Println("从b中取到了", x)
	// }()
	b <- 10
	fmt.Println("10发送到通道b了")
	fmt.Println(b)
	// wg.Wait()
}
