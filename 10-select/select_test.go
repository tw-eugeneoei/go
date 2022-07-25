package select_tut

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// * v1
// func TestRacer(t *testing.T) {
// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(1 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	slowUrl := "https://www.facebook.com/"
// 	fastUrl := "https://quii.gitbook.io/learn-go-with-tests/"

// 	expected := fastUrl
// 	result := Racer(slowUrl, fastUrl)

// 	if result != expected {
// 		t.Errorf("expected %s, result is %s", expected, result)
// 	}

// 	slowServer.Close()
// 	fastServer.Close()
// }

// * v2
func TestRacer(t *testing.T) {

	t.Run("should return fast url", func(t *testing.T) {
		slowServer := createFakeServer(20 * time.Millisecond)
		fastServer := createFakeServer(1 * time.Millisecond)

		// * By prefixing a function call with defer it will now call that function at the end of the containing function.
		// * You want this to execute at the end of the function, but keep the instruction near where you created the server
		// * for the benefit of future readers of the code.
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		result, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Errorf("expected no error but received error %v", err)
		}

		if result != expected {
			t.Errorf("expected %s, result is %s", expected, result)
		}

	})

	t.Run("should return an error if server does not respond within 10s", func(t *testing.T) {
		server := createFakeServer(25 * time.Millisecond)
		timeout := 20 * time.Millisecond

		defer server.Close()

		url := server.URL

		_, err := ConfigurableRacer(url, url, timeout)

		if err == nil {
			t.Errorf("expected err but did not throw one")
		}
	})

}

func createFakeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
