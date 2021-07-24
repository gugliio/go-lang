package main

import (
	"fmt"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	Workerpool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerpool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerpool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d ended with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with id %d stoped\n", w.Id)
			}
		}
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func StopWorker(w Worker) {
	go func() {
		w.QuitChan <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		Workerpool: worker,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.Workerpool
				workerJobQueue <- job
			}()
		}
	}
}
