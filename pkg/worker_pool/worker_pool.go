package worker_pool

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker int
	Tasks     chan func()
}

func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker: maxWorker,
		Tasks:     make(chan func()),
	}

	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(task func()) {
	wp.Tasks <- task
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func() {
			for task := range wp.Tasks {
				task()
			}
		}()
	}
}
