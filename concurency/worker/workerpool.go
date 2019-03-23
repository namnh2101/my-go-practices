package workerpool

import (
	"context"
	"sync"
)

// TaskResult represent result of task
type TaskResult struct {
	Result interface{}
	Err    error
}

// Task represent a task
type Task struct {
	ctx      context.Context
	executor func(context.Context) (interface{}, error)
	future   chan *TaskResult
}

// NewTask create new task
func NewTask(ctx context.Context, executor func(context.Context) (interface{}, error)) *Task {
	return &Task{
		ctx:      ctx,
		executor: executor,
		future:   make(chan *TaskResult, 1),
	}
}

// Execute task
func (t *Task) Execute() {
	var result interface{}
	var err error

	if t.executor != nil {
		result, err = t.executor(t.ctx)
	}

	t.future <- &TaskResult{Result: result, Err: err}
}

// Result pushed via channel
func (t *Task) Result() <-chan *TaskResult {
	return t.future
}

// Pool of worker
type Pool struct {
	ctx          context.Context
	cancel       context.CancelFunc
	numberWorker int
	wg           sync.WaitGroup
	ch           chan *Task
	chBatch      chan []*Task
}

// NewPool create new worker pool
func NewPool(ctx context.Context, numberWorker int) (p *Pool) {
	if numberWorker <= 0 {
		numberWorker = 1
	}

	if ctx == nil {
		ctx = context.Background()
	}

	p = &Pool{
		numberWorker: numberWorker,
		ch:           make(chan *Task, numberWorker),
		chBatch:      make(chan []*Task, numberWorker),
	}
	p.ctx, p.cancel = context.WithCancel(ctx)

	return
}

// Start workers
func (p *Pool) Start() {
	p.wg.Add(p.numberWorker + 1)

	// single task
	for i := 0; i < p.numberWorker; i++ {
		go p.worker()
	}

	// batched task
	go func() {
		defer p.wg.Done()

		var (
			ctx     = p.ctx
			chBatch = p.chBatch
			tasks   []*Task
			task    *Task
			ok      bool
		)

		for {
			select {

			case tasks, ok = <-chBatch:
				if !ok {
					return
				}

				if len(tasks) > 0 {
					for _, task = range tasks {
						if p.ch != nil && task != nil {
							select {
							case p.ch <- task:

							case <-ctx.Done():
								return
							}
						}
					}
				}

			case <-ctx.Done():
				return
			}
		}
	}()
}

// Stop worker. Wait all task done.
func (p *Pool) Stop() {
	// cancel context
	p.cancel()

	// wait child workers
	p.wg.Wait()
}

// Do a task
func (p *Pool) Do(t *Task) {
	if p.ch != nil && t != nil {
		select {
		case <-p.ctx.Done():
		case p.ch <- t:
		}
	}
}

// Batch execute batch job
func (p *Pool) Batch(tasks []*Task) {
	p.chBatch <- tasks
}

// Execute a task
func (p *Pool) Execute(exec func(context.Context) (interface{}, error)) (t *Task) {
	return p.ExecuteWithCtx(p.ctx, exec)
}

// ExecuteWithCtx a task with custom context
func (p *Pool) ExecuteWithCtx(ctx context.Context, exec func(context.Context) (interface{}, error)) (t *Task) {
	if ctx == nil {
		ctx = p.ctx
	}
	t = NewTask(ctx, exec)
	p.Do(t)
	return
}

func (p *Pool) worker() {
	defer p.wg.Done()

	var task *Task

	for {
		select {
		case <-p.ctx.Done():
			return

		case task = <-p.ch:
			if task != nil {
				task.Execute()
			}
		}
	}
}
