package synctut

import "sync"

type Counter struct {
	// * A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu    sync.Mutex
	count int
}

func (c *Counter) inc() {
	// * What this means is any goroutine calling "inc" will acquire the lock on Counter if they are first.
	// * All the other goroutines will have to wait for it to be Unlocked before getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += 1
}

func (c *Counter) value() int {
	return c.count
}
