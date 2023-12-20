package controllers

import (
	"sync"

	"github.com/BerkBugur/Go-Project/initializers"
	"github.com/BerkBugur/Go-Project/models"
)

type TaskJob struct {
	Task       models.Task
	Result     chan error
	TaskResult chan models.Task // Updated Task
}

type WorkerPool struct {
	workerCount int
	jobQueue    chan TaskJob
	wg          sync.WaitGroup
}

var pool *WorkerPool

func init() {
	pool = NewWorkerPool(5) // Worker Thread number
	pool.Start()
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		jobQueue:    make(chan TaskJob),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	for job := range wp.jobQueue {
		result := initializers.DB.Create(&job.Task)
		job.Result <- result.Error
		if result.Error == nil {
			job.TaskResult <- job.Task // Update task and return back
		}
	}
	wp.wg.Done()
}
