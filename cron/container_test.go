package cron

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCronStopDuringStart(t *testing.T) {
	const iterations = 100
	for i := 0; i < iterations; i++ {
		container := NewContainer()

		var startStop sync.WaitGroup
		startStop.Add(2)
		go func() {
			defer startStop.Done()
			container.Start()
		}()
		go func() {
			defer startStop.Done()
			container.Stop()
		}()
		startStop.Wait()

		// Stop before Start is permitted to be a no-op. If Start won the race,
		// ensure the started workers are stopped before this iteration ends.
		container.Stop()
	}
}

func TestCronConcurrentStartAndStop(t *testing.T) {
	container := NewContainer()
	container.Add(NewTask("start-stop", time.Millisecond, func(context.Context) {}))

	container.Start()
	var stops sync.WaitGroup
	for i := 0; i < 10; i++ {
		stops.Add(1)
		go func() {
			defer stops.Done()
			container.Stop()
		}()
	}
	stops.Wait()
}

func TestCron(t *testing.T) {
	container := NewContainer()

	counter := 0
	container.Add(NewTask("test task", 20*time.Millisecond, func(ctx context.Context) {
		counter++
	}))

	container.Start()
	time.Sleep(2 * time.Second)

	container.Stop()

	t.Log(counter)

	assert.Greater(t, counter, 90)
	assert.Less(t, counter, 110)
}
