package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

// ResponseError result from http.get for an specific url
type ResponseError struct {
	Resp *http.Response
	Err  error
}

// IxResponseError result from http.get for an specific url
type IxResponseError struct {
	Ix   int
	Resp *http.Response
	Err  error
}

// GetAllUrls get all urls and retuns when all of them have finished
func GetAllUrls(urls []string) []ResponseError {
	dones := make(chan IxResponseError)
	for ix, url := range urls {
		go func(i int, u string, done chan IxResponseError) {
			fmt.Printf("getting: index: %d, %s\n", i, urls[i])
			r, err := http.Get(u)
			fmt.Printf("done with: %d, %s\n", i, urls[i])
			done <- IxResponseError{Ix: i, Resp: r, Err: err}

		}(ix, url, dones)
	}
	ret := make([]ResponseError, len(urls))
	for i := 0; i < len(urls); i++ {
		r := <-dones
		ret[r.Ix] = ResponseError{Resp: r.Resp, Err: r.Err}
	}

	return ret
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	urls := []string{"https://google.com", "http://www.degconnect.com", "http://elpais.com", "https://yahoo.es", "https://yahoo.com",
		"http://www.degconnect.com"}

	start := time.Now()
	results := GetAllUrls(urls)
	t := time.Now()
	elapsed := t.Sub(start)

	for ix, r := range results {
		if r.Err == nil {
			body, _ := ioutil.ReadAll(r.Resp.Body)
			fmt.Printf("result of %s is [%v]\n", urls[ix], string(body)[:15])
		} else {
			fmt.Printf("result of %s is [%v]\n", urls[ix], r.Err)
		}

	}

	fmt.Printf("Elapsed: %v\n", elapsed)

}
