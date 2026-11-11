package cron

import (
	"context"
	"sync"
	"time"
)

func NewContainer() *Container {
	return &Container{}
}

type Container struct {
	mu      sync.Mutex
	tasks   []*Task
	wg      sync.WaitGroup
	stop    context.CancelFunc
	started bool
}

func (c *Container) Add(task *Task) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tasks = append(c.tasks, task)
}

func (c *Container) Start() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return
	}
	c.started = true
	ctx, cancel := context.WithCancel(context.Background())
	c.stop = cancel
	c.wg = sync.WaitGroup{}
	for _, task := range c.tasks {
		c.wg.Add(1)
		go taskWrapper(ctx, &c.wg, task)
	}
}

func (c *Container) Stop() {
	c.mu.Lock()
	cancel := c.stop
	c.mu.Unlock()

	if cancel != nil {
		cancel()
	}
	c.wg.Wait()
	c.mu.Lock()
	c.started = false
	c.stop = nil
	c.mu.Unlock()
}

func taskWrapper(ctx context.Context, wg *sync.WaitGroup, task *Task) {
	defer wg.Done()

	ticker := time.NewTicker(task.Frequency)
	defer ticker.Stop()
	done := ctx.Done()

	for {
		select {
		case <-ticker.C:
			task.F(ctx)
		case <-done:
			return
		}
	}
}
