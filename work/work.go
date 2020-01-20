/*Package work handles goroutine addational funcs*/
package work

import (
	"sync"
	"time"
)

// WaitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

/*TimeoutFunc executes a callback function every second*/
func TimeoutFunc(stop chan struct{}, callback func() error, timeout time.Duration) error {
	//optional with ttl
	ttl := make(<-chan time.Time, 1) // placeholder for timer, 0 run forever
	if timeout > 0 {
		ttl = time.After(timeout)
	}

	for {
		select {
		case <-stop:
			return nil
		case <-ttl:
			return nil
		case <-time.After(1 * time.Second):
			if err := callback(); err != nil {
				return err
			}
		}
	}
}
