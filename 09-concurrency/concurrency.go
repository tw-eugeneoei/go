package concurrency

type WebsiteChecker func(string) bool

// * decalring anonymous fields in struct ie struct type without declaring any field names
// * field names default to type declaration
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// * WRITING data into channel
			// * note that this is a blocking call
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// * READING data from channel
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
