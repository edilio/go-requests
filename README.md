# go-requests
Go snippet to get a bunch of urls as inputs and return all the responses sorted in the same order they are in input

## Similar to one of the uses of grequests in Python
https://github.com/kennethreitz/grequests


I haven't seen so much simplecity in any other programming language or library for getting this done

```python
import grequests

urls = [
    'http://www.heroku.com',
    'http://python-tablib.org',
    'http://httpbin.org',
    'http://python-requests.org',
    'http://fakedomain/',
    'http://kennethreitz.com'
]

rs = (grequests.get(u) for u in urls)

for r in grequests.map(rs):
    process(r)
```

## Now in Golang
There is a lot of concepts in golang for concurrency but I missed the simplicity in grequests so I created this snippet as a reminder

```go
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
``` 

It can be used later as

```go
    urls := []string{"https://google.com", "http://www.degconnect.com", "http://elpais.com", "https://yahoo.es", "https://yahoo.com",
		"http://www.degconnect.com"}

	results := GetAllUrls(urls)
	
	for ix, r := range results {
		if r.Err == nil {
			body, _ := ioutil.ReadAll(r.Resp.Body)
			fmt.Printf("result of %s is [%v]\n", urls[ix], string(body)[:15])
		} else {
			fmt.Printf("result of %s is [%v]\n", urls[ix], r.Err)
		}

	}
```


