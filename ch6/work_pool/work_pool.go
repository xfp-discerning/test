package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	//开启三个gorutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}
	//5个任务
	func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
	}()
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Println(<-result)
	}
}
