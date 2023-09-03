package semaphore

import (
	"log"
	"sync"
	"testing"
	"time"
)

func process(id int, calls *int, wg *sync.WaitGroup) {
	log.Printf("[process %d]: running", id)
	*calls++
	time.Sleep(5 * time.Second)
	wg.Done()
}

func TestSemaphore(t *testing.T) {
	maxConcurrency := 5
	processesNum := 10
	calls := 0

	t.Logf("starting semaphore with %d max_concurrency and %d total tasks", maxConcurrency, processesNum)

	var wg sync.WaitGroup

	semaphore := NewSemaphore(5)
	for i := 0; i < processesNum; i++ {
		wg.Add(1)
		semaphore.Acquire()
		go func(processID int) {
			defer semaphore.Release()
			process(processID, &calls, &wg)
			log.Printf("[process %d]: finished", processID)
		}(i + 1)
	}

	wg.Wait()
	if calls != processesNum {
		t.Errorf("expected %d tasks to be executed, but got %d", processesNum, calls)
	}
}
