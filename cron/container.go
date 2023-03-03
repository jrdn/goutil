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
	tasks []*Task
	wg    sync.WaitGroup
}

func (c *Container) Add(task *Task) {
	c.tasks = append(c.tasks, task)
}

func (c *Container) Start() {
	c.wg = sync.WaitGroup{}
	for _, task := range c.tasks {
		c.wg.Add(1)
		go taskWrapper(context.Background(), &c.wg, task)
	}
}

func (c *Container) Stop() {
	for _, t := range c.tasks {
		t.cancel()
	}
	c.wg.Wait()
}

func taskWrapper(ctx context.Context, wg *sync.WaitGroup, task *Task) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(ctx)
	task.cancel = cancel

	ticker := time.NewTicker(task.Frequency)
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
