package main

import "sync"

type CustomCounter struct {
	mu    sync.Mutex
	value int64
}

func (c *CustomCounter) Inc(val int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += val
}

func (c *CustomCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
