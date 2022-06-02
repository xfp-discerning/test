package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

func creat(jobchan chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63n(100000)
		var jobs = &job{
			value: x,
		}
		jobchan <- jobs
		time.Sleep(time.Millisecond * 500)
	}
}
func sums(jobchan <-chan *job, resultchan chan<- *result) {
	defer wg.Done()
	for {
		jobs := <-jobchan
		sum := int64(0)
		n := jobs.value
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		newreualt := &result{
			job:    jobs,
			result: sum,
		}
		resultchan <- newreualt
	}
}

func main() {
	jobchan := make(chan *job, 50)
	resultchan := make(chan *result, 100)
	wg.Add(1)
	go creat(jobchan)
	wg.Add(24)
	for i := 1; i <= 24; i++ {
		go sums(jobchan, resultchan)
	}
	for reslut := range resultchan {
		fmt.Printf("value:%d, sum:%d\n", reslut.job.value, reslut.result)
	}
	wg.Wait()
}
