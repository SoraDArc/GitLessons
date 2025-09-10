package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, taskSet chan<- int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("День у рабочего", n, "закончился!")
			return
		default:
			fmt.Println("Рабочий", n, "начал работу!")
			time.Sleep(1 * time.Second)
			fmt.Println("Рабочий", n, "закончил работу!", "Сделал", n*10, "задач")
			taskSet <- n * 10
		}
	}
}

func WorkerPool(ctx context.Context, workerCnt int) <-chan int {
	taskSetPoint := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= workerCnt; i++ {
		wg.Add(1)
		go worker(ctx, taskSetPoint, &wg, i)
	}

	go func() {
		defer close(taskSetPoint)
		wg.Wait()
	}()

	return taskSetPoint
}
