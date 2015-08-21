package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/satoshun/goworker"
)

func main() {
	workerSize := 10
	service := goworker.NewService(workerSize)
	go service.Start()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		fmt.Printf("Generate Task: %d\n", i)
		func(index int) {
			service.Run(func() {
				time.Sleep(time.Second * 1)
				fmt.Printf("Complete Task: %d\n", index)
				wg.Done()
			})
		}(i)
	}

	wg.Wait()
}
