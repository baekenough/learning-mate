package main

import (
	"fmt"
	"sync"
	"time"
)

type CrawlResult struct {
	url  string
	body string
}

type SafeCrawlStore struct {
	mu     sync.Mutex
	result map[string]string
}

func worker(id int, jobs chan string, store *SafeCrawlStore, wg *sync.WaitGroup) string {
	url := <-jobs
	store.mu.Lock()
	store.result[url] = "Crawled " + url
	defer store.mu.Unlock()
	wg.Done()
	return store.result[url]
}

func main() {
	var target = []string{"naver.com", "google.com", "yahoo.com", "hi.com", "notion.io"}
	jobs := make(chan string, len(target))
	for i := 0; i < len(target); i++ {
		jobs <- target[i]
	}
	instance := SafeCrawlStore{result: make(map[string]string)}

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		tick := time.NewTicker(50 * time.Millisecond)
		wg.Add(1)
		go worker(i, jobs, &instance, &wg)
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		select {
		case <-tick.C:
			fmt.Println(len(instance.result))
		case <-done:
			tick.Stop()
			fmt.Println("All workers done.")
			return
		default:
			time.Sleep(1 * time.Second)
			for result := range instance.result {
				fmt.Println(result)
			}
		}
	}

}
