package mocking

import (
	"fmt"
	"io"
	"time"
)

const start = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type CustomSleeper struct{}

func (d *CustomSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)

	// for i := start; i > 0; i-- {
	// 	sleeper.Sleep()
	// }

	// for i := start; i > 0; i-- {
	// 	fmt.Fprintln(writer, i)
	// }

	// fmt.Fprint(writer, finalWord)
}
