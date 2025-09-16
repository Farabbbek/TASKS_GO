package main

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type StatusResult struct {
	URL  string
	Code int   // 0, если не удалось получить ответ
	Err  error // nil, если ответ получен
}

func FetchStatuses(
	urls []string,
	concurrency int, // общий лимит одновременных запросов
	timeout time.Duration,
) []StatusResult {
	if concurrency < 1 {
		concurrency = 1
	}
	var wait sync.WaitGroup
	res := make([]StatusResult, len(urls))
	sem := make(chan struct{}, concurrency)

	for i, url := range urls {
		wait.Add(1)
		go func(i int, url string) {
			defer wait.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				res[i] = StatusResult{URL: url, Code: 0, Err: err}
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				res[i] = StatusResult{URL: url, Code: 0, Err: err}
				return
			}
			res[i] = StatusResult{URL: url, Code: resp.StatusCode, Err: nil}
			resp.Body.Close()
		}(i, url)
	}
	wait.Wait()
	return res
}
