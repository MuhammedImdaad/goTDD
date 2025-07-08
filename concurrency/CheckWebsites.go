package concurrency

type WebsiteChecker func(string) bool

type result struct {
	s string
	b bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	ch := make(chan result) // By default, sends and receives block until the other side is ready

	for _, url := range urls {
		go func() { // the only way to start a `goroutine` is to put go in front of a function call
			ch <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <- ch
		results[r.s] = r.b
	}

	return results
}
