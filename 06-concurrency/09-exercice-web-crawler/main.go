package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Creates a mutex we can use to control the urls
var fetchedUrls = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("Url load is in progress...")

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	// First lock the mutex to avoid multiple check at the same time
	fetchedUrls.Lock()

	// Check if the url is already in our map of fetched urls
	if _, ok := fetchedUrls.m[url]; ok {
		fetchedUrls.Unlock()
		fmt.Printf("<- done with %v, url already fetched! \n", url)
		return
	}

	// It is not. -> we can flag the url as loading and unlock the mutex (flag
	// first and unlock after)
	fetchedUrls.m[url] = loading
	fetchedUrls.Unlock()

	// Fetches the url
	body, urls, err := fetcher.Fetch(url)

	// Update the falg on the url to set it explicitely to "fetched"
	// Start by Locking then update and then unlock
	fetchedUrls.Lock()
	fetchedUrls.m[url] = err
	fetchedUrls.Unlock()

	// checks that the fetched did not fail before continuing
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	// Creates a channel to pass all the crawl result along and know when we are
	// done. (could we eventually use a buffered channel? looks like it but not
	// sure about possible side effects...)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		crawlFoundUrl := func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true // use the channel to be able to "block" below until until
			// the crawl has been done. it will help us synchronize the end of the
			// process
		}

		go crawlFoundUrl(u)
	}

	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done // wait until the other side is ready
	}
	fmt.Printf("<- Done with %v\n", url)
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
	fmt.Println("Fetching stats\n--------------")
	for url, err := range fetchedUrls.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
