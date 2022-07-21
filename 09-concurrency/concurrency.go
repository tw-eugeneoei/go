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
			// * sending a result struct for each call to wc to the resultChannel with a send statement
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// * we're using a receive expression, which assigns a value received from a channel to a variable
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
