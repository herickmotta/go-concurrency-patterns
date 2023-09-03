package workerpool

import "log"

/*
	A worker pool is a collection of threads(goroutines) that are waiting for tasks to be assigned to them.
	Creating a goroutine whenever a task needs to be executed can be dangerous, leading to CPU or MEMORY issues.
	Managing goroutines with a thread pool can be a good way to get done tasks individually.

	In the code below, we create a simple struct of a thread pool, that receives and runs tasks through a channel.
*/

type Task struct {
	ID  int
	Run func()
}

type WorkerPool interface {
	Run()
	AddTask(task Task)
}

type workerPool struct {
	maxWorkers      int
	queuedTasksChan chan Task
}

func NewWorkerPool(maxWorkers int) WorkerPool {
	return &workerPool{
		maxWorkers:      maxWorkers,
		queuedTasksChan: make(chan Task),
	}
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorkers; i++ {
		go func(workerID int) {
			log.Printf("[worker %d]: starting", workerID)
			for task := range wp.queuedTasksChan {
				log.Printf("[worker %d]: processing task %d", workerID, task.ID)
				task.Run()
				log.Printf("[worker %d]: task %d finished", workerID, task.ID)
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task Task) {
	wp.queuedTasksChan <- task
}
