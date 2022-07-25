package select_tut

import (
	"fmt"
	"net/http"
	"time"
)

// * v1
// func Racer(urlOne, urlTwo string) (winner string) {
// 	urlOneDuration := calculateResponseTime(urlOne)
// 	urlTwoDuration := calculateResponseTime(urlTwo)

// 	if urlOneDuration < urlTwoDuration {
// 		return urlOne
// 	}

// 	return urlTwo
// }

// func calculateResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	duration := time.Since(start)
// 	return duration
// }

const tenSecondTimeout = 10 * time.Second

// * v2
func Racer(urlOne, urlTwo string) (winner string, err error) {
	return ConfigurableRacer(urlOne, urlTwo, tenSecondTimeout)
}

func ConfigurableRacer(urlOne, urlTwo string, timeout time.Duration) (winner string, err error) {
	// * select lets you wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for response")
	}
}

// * Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective
// * so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
func ping(url string) chan struct{} {
	// * Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
	// * For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func Racer(a, b string) (winner string) {
// 	select {
// 	case <-ping(a):
// 		return a
// 	case <-ping(b):
// 		return b
// 	}
// }

// func ping(url string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		http.Get(url)
// 		close(ch)
// 	}()
// 	return ch
// }
