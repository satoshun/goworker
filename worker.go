package goworker

// Task is be executed by a worker.
type Task interface {
	Run()
}

// Runnable is be executed by a worker.
type Runnable func()

type runnable struct {
	task Runnable
}

func (r runnable) Run() {
	r.task()
}

type worker struct {
	channel     chan Task
	workerQueue chan chan Task
}

func (w *worker) start() {
	go func() {
		for {
			w.workerQueue <- w.channel

			select {
			case task := <-w.channel:
				task.Run()
			}
		}
	}()
}

// WorkerService is a goroutine pool.
type WorkerService struct {
	taskQueue   chan Task
	workerQueue chan chan Task
	workerSize  int
}

// RunTask run specific task.
func (s *WorkerService) RunTask(task Task) {
	s.taskQueue <- task
}

// Run execute specific task.
func (s *WorkerService) Run(task Runnable) {
	s.taskQueue <- runnable{task: task}
}

// Start is entrypoint.
func (s *WorkerService) Start() {
	for i := 0; i < s.workerSize; i++ {
		w := worker{
			channel:     make(chan Task),
			workerQueue: s.workerQueue,
		}
		w.start()
	}

	for {
		select {
		case task := <-s.taskQueue:
			go func(t Task) {
				queue := <-s.workerQueue
				queue <- t
			}(task)
		}
	}
}

// NewService is factory method of WorkerService.
func NewService(workerSize int) *WorkerService {
	return &WorkerService{
		taskQueue:   make(chan Task),
		workerQueue: make(chan chan Task, workerSize),
		workerSize:  workerSize,
	}
}
