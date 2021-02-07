// Exercise: Web Crawler
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	mu sync.Mutex
	m  map[string]bool
}

type Counter struct {
	mu     sync.Mutex
	number int
}

func (c *Counter) Add(v int) {
	c.mu.Lock()
	c.number += v
	c.mu.Unlock()
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.number
}

func Crawl(url string, depth int, fetcher Fetcher) {
	visited := Cache{m: make(map[string]bool)}
	numWorker := Counter{number: 1}
	crawl(url, depth, fetcher, &visited, &numWorker)
	for numWorker.Get() > 0 {
	}
}

func crawl(url string, depth int, fetcher Fetcher, visited *Cache, numWorker *Counter) {
	defer numWorker.Add(-1)

	visited.mu.Lock()
	exists := visited.m[url]
	visited.m[url] = true
	visited.mu.Unlock()

	if exists || depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	numWorker.Add(len(urls))
	for _, u := range urls {
		go crawl(u, depth-1, fetcher, visited, numWorker)
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
