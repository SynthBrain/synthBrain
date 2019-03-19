package samples

import (
	"fmt"
	"time"
)

func Do(workers int, f func() error) (int, error) {
	ch := make(chan error, workers)

	for i := 0; i < workers; i++ {
		go func(n int) {
			fmt.Printf("worker %d", n)
			ch <- f()
		}(i)
	}

	for i := 0; i < workers; i++ {
		select {
		case err := <-ch:
			if err != nil {
				return 0, nil
			}
		case <-time.After(time.Second):
			return 0, fmt.Errorf("timeOut")
		}
	}
	return workers, nil
}
