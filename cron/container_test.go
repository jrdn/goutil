package cron

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCron(t *testing.T) {
	container := NewContainer()

	counter := 0
	container.Add(NewTask("test task", 5*time.Millisecond, func(ctx context.Context) {
		counter++
	}))

	container.Start()
	time.Sleep(1 * time.Second)

	container.Stop()

	t.Log(counter)

	assert.Greater(t, counter, 175)
}
