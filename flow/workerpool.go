package flow

import "sync"

type WorkerPool struct {
	c  chan func()
	wg sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	p := &WorkerPool{c: make(chan func())}
	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			for task := range p.c {
				task()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *WorkerPool) Run(task func()) {
	p.c <- task
}

func (p *WorkerPool) Wait() {
	close(p.c)
	p.wg.Wait()
}
