package concurrency

import (
	"testing"
	"time"
)

// // * normal
// func mockWebsiteChecker(url string) bool {
// 	if url == "waat://furhurterwe.geds" {
// 		return false
// 	}
// 	return true
// }

// func TestCheckWebsites(t *testing.T) {
// 	websites := []string{
// 		"http://google.com",
// 		"http://blog.gypsydave5.com",
// 		"waat://furhurterwe.geds",
// 	}

// 	want := map[string]bool{
// 		"http://google.com":          true,
// 		"http://blog.gypsydave5.com": true,
// 		"waat://furhurterwe.geds":    false,
// 	}

// 	got := CheckWebsites(mockWebsiteChecker, websites)

// 	fmt.Println(">>>>>", reflect.DeepEqual(want, got))

// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}
// }

// * slow stub to simulate slow checking
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	size := 100
	urls := make([]string, 100)
	for i := 0; i < size; i++ {
		urls[i] = "a url"
	}

	// urls := []string{}
	// for i := 0; i < size; i++ {
	// 	urls = append(urls, "a url")
	// }

	b.ResetTimer()
	// * The b.N is a variable to be adjusted by Go.
	// * Go chooses it after doing some small adjustments based on its internally defined rules.
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
