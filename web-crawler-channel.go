// Exercise: Web Crawler using channel
package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Pair struct {
	urls  []string
	depth int
}

func crawRoutine(url string, depth int, ch chan Pair, fetcher Fetcher) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- Pair{[]string{}, 0}
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	if depth < 1 {
		ch <- Pair{[]string{}, 0}
	} else {
		ch <- Pair{urls, depth - 1}
	}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	cache := make(map[string]bool)
	ch := make(chan Pair, 10)
	ch <- Pair{[]string{url}, depth}
	n := 1
	for p := range ch {
		for _, u := range p.urls {
			if !cache[u] {
				cache[u] = true
				n += 1
				go crawRoutine(u, p.depth, ch, fetcher)
			}
		}
		n -= 1
		if n == 0 {
			break
		}
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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
