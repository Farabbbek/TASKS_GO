package main

import (
	"sync"
)

type Job func() error

func RunAll(jobs []Job) error {
	var wait sync.WaitGroup
	var erer error
	var once sync.Once
	wait.Add(len(jobs))
	for _, job := range jobs {
		go func(j Job) {
			defer wait.Done()
			if err := j(); err != nil {
				once.Do(func() { erer = err })
			}
		}(job)
	}
	wait.Wait()
	return erer
}
