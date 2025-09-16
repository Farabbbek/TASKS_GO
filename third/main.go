package main

import (
	"sync"
)

type MapFunc[T any, R any] func(T) R

func ParallelMap[T any, R any](in []T, fn MapFunc[T, R], workers int) []R {
	out := make([]R, len(in))
	if workers < 1 {
		workers = 1
	}
	var wait sync.WaitGroup
	tasks := make(chan struct {
		a int
		b T
	}, len(in))
	for w := 0; w < workers; w++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for task := range tasks {
				out[task.a] = fn(task.b)
			}
		}()
	}

	for i, v := range in {
		tasks <- struct {
			a int
			b T
		}{i, v}
	}

	close(tasks)
	wait.Wait()
	return out
}
