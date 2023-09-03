package workerpool

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	workersNum := 3
	tasksNum := 10

	t.Logf("starting worker pool with %d workers and %d total tasks", workersNum, tasksNum)

	wp := NewWorkerPool(workersNum)
	wp.Run()

	routinesNum := runtime.NumGoroutine()
	tasksDone := 0

	var wg sync.WaitGroup
	for i := 0; i < tasksNum; i++ {
		wg.Add(1)
		wp.AddTask(Task{
			ID: i + 1,
			Run: func() {
				defer wg.Done()
				tasksDone++
				time.Sleep(2 * time.Second)
			},
		})
	}

	wg.Wait()

	if tasksDone != tasksNum {
		t.Errorf("expected %d tasks to be executed, but got %d", tasksNum, tasksDone)
	}

	if routinesNum < workersNum {
		t.Errorf("expected at least %d goroutines to be running, but got %d", workersNum, routinesNum)
	}
}
