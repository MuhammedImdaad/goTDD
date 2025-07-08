package concurrency

import "sync"

type Counter struct {
	count int
	mu sync.Mutex // A Mutex is a MUTual EXclusion lock.
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c* Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c* Counter) Value() int {
	return c.count
}

