package timeout_test

import (
	"atomicgo.dev/timeout"
	"fmt"
	"time"
)

func Example_timeoutReached() {
	res, err := timeout.Execute(time.Second*1, func() (string, error) {
		time.Sleep(time.Second * 2)
		return "Hello, World!", nil
	})

	fmt.Println(res, err)
	// Output:
	//  timeout reached
}

func Example_timeoutNotReached() {
	res, err := timeout.Execute(time.Second*2, func() (string, error) {
		time.Sleep(time.Second * 1)
		return "Hello, World!", nil
	})

	fmt.Println(res, err)
	// Output:
	// Hello, World! <nil>
}

func Example_timeoutWithError() {
	res, err := timeout.Execute(time.Second*2, func() (string, error) {
		time.Sleep(time.Second * 1)
		return "", fmt.Errorf("some error")
	})

	fmt.Println(res, err)
	// Output:
	//  some error
}
