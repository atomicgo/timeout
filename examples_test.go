package timeout_test

import (
	"atomicgo.dev/template"
	"fmt"
)

func Example_demo() {
	fmt.Println(timeout.HelloWorld())
	// Output: Hello, World!
}

func ExampleHelloWorld() {
	fmt.Println(timeout.HelloWorld())
	// Output: Hello, World!
}
