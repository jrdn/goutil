package cron

import (
	"context"
	"time"
)

type TaskFunc func(ctx context.Context)

type Task struct {
	Name      string
	Frequency time.Duration
	F         func(ctx context.Context)

	cancel context.CancelFunc
}

func NewTask(name string, freq time.Duration, f TaskFunc) *Task {
	return &Task{
		Name:      name,
		Frequency: freq,
		F:         f,
	}
}
