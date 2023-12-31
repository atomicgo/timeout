package timeout

import (
	"errors"
	"time"
)

// Function is a function that returns a generic value and an error.
type Function[T any] func() (T, error)

// ErrTimeoutReached is returned when the timeout is reached.
var ErrTimeoutReached = errors.New("timeout reached")

// Execute executes a function and returns the result or an error if the timeout is reached.
// If the timeout is reached, the Function will be interrupted.
// If the Function returns an error, the error will be returned.
func Execute[T any](duration time.Duration, fn Function[T]) (T, error) {
	resultChannel := make(chan T, 1)
	errorChannel := make(chan error, 1)

	go func() {
		result, err := fn()
		if err != nil {
			errorChannel <- err
			return
		}
		resultChannel <- result
	}()

	select {
	case res := <-resultChannel:
		return res, nil
	case err := <-errorChannel:
		return zeroValue[T](), err
	case <-time.After(duration):
		return zeroValue[T](), ErrTimeoutReached
	}
}

func zeroValue[T any]() T {
	var v T
	return v
}
