package utils

import "time"

func Timeout(fn func(), duration time.Duration) func() {
	shouldGo := true

	go func() {
		time.Sleep(duration)

		if shouldGo {
			fn()
		}
	}()

	return func() {
		shouldGo = false
	}
}
