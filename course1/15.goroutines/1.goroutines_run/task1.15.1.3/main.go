package main

import (
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID     int
	Data   string
	Status bool
}

type WorkerPool struct {
	Queue        chan *Task
	WorkersCount int
	Result       chan *Task
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{Queue: make(chan *Task),
		WorkersCount: workers,
		Result:       make(chan *Task)}
}

func (wp *WorkerPool) Start() {
	go func() {
		for task := range wp.Queue {
			wp.worker(task)
		}
	}()
}

func (wp *WorkerPool) worker(task *Task) {
	log.Printf("Executing task %d: %s\n", task.ID, task.Data)
	task.Status = true
	wp.Result <- task
	wp.WorkersCount--

}

func (wp *WorkerPool) AddTask(task *Task) {
	wp.Queue <- task
	wp.WorkersCount++

}

// Wait - waiting for all tasks to complete in seconds
func (wp *WorkerPool) Wait(timer int) {
	time.Sleep(time.Duration(timer) * time.Second)
}

func main() {
	pool := NewWorkerPool(5)
	fmt.Println(pool)
	pool.Start()

	tasks := []*Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
	}

	go func() {
		for _, task := range tasks {
			pool.AddTask(task)
		}
	}()

	go func() {
		for v := range pool.Result {
			log.Printf("Task %d completed: %s status: %v\n", v.ID, v.Data, v.Status)
		}
	}()
	// Waiting for all tasks to complete
	pool.Wait(1)
	close(pool.Queue)
	close(pool.Result)
}
