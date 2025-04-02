package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Number int
	Name   string
	Delay  time.Duration
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker %d started job %d\n", w.Id, job.Number)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker %d finished job %d with result %d\n", w.Id, job.Number, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker %d stopped\n", w.Id)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}

func (d *Dispatcher) Dispatch() {
	for job := range d.JobQueue {
		go func() {
			workerJobQueue := <-d.WorkerPool
			workerJobQueue <- job
		}()
	}
}

func (d *Dispatcher) Run() {
	for i := range d.MaxWorkers {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := Job{
		Number: value,
		Name:   name,
		Delay:  delay,
	}

	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	const maxWorkers = 4
	const maxQueue = 20
	const port = ":8081"

	jobQueue := make(chan Job, maxQueue)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
