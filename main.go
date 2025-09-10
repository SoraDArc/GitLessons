package main

import (
	"context"
	"fmt"
	"project/worker"
	"sync"
	"time"
)

func main() {
	workerContext, workerCancel := context.WithCancel(context.Background())

	timeInit := time.Now()
	taskSetPoint := worker.WorkerPool(workerContext, 100)
	time.Sleep(3 * time.Second)
	workerCancel()

	var mtx sync.Mutex
	var wg sync.WaitGroup
	taskResult := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range taskSetPoint {
			mtx.Lock()
			taskResult += val
			mtx.Unlock()
		}
	}()
	wg.Wait()

	fmt.Println("Программа завершилась! Итого работ выполнено:", taskResult)
	fmt.Println("Программа работала:", time.Since(timeInit))
}
