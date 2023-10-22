<h1 align="center">AtomicGo | timeout</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Ftimeout&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/timeout/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/timeout?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/timeout" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/timeout/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/timeout" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/timeout?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/timeout">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-3-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/timeout" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/timeout?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/timeout#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/timeout</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# timeout

```go
import "atomicgo.dev/timeout"
```

Package timeout provides a simple way to add timeouts to your Go code. The Execute function is blocking, so you can use it as a drop\-in replacement for your function calls. Timeouts are often useful in scripts, where you want to limit the execution time of a function.

<details><summary>Example (Timeout Not Reached)</summary>
<p>



```go
package main

import (
	"atomicgo.dev/timeout"
	"fmt"
	"time"
)

func main() {
	res, err := timeout.Execute(time.Second*2, func() (string, error) {
		time.Sleep(time.Second * 1)
		return "Hello, World!", nil
	})

	fmt.Println(res, err)
}
```

#### Output

```
Hello, World! <nil>
```

</p>
</details>

<details><summary>Example (Timeout Reached)</summary>
<p>



```go
package main

import (
	"atomicgo.dev/timeout"
	"fmt"
	"time"
)

func main() {
	res, err := timeout.Execute(time.Second*1, func() (string, error) {
		time.Sleep(time.Second * 2)
		return "Hello, World!", nil
	})

	fmt.Println(res, err)
}
```

#### Output

```
timeout reached
```

</p>
</details>

<details><summary>Example (Timeout With Error)</summary>
<p>



```go
package main

import (
	"atomicgo.dev/timeout"
	"errors"
	"fmt"
	"time"
)

func main() {
	res, err := timeout.Execute(time.Second*2, func() (string, error) {
		time.Sleep(time.Second * 1)
		return "", errors.New("some error") // nolint: goerr113
	})

	fmt.Println(res, err)
}
```

#### Output

```
some error
```

</p>
</details>

## Index

- [Variables](<#variables>)
- [func Execute\[T any\]\(duration time.Duration, fn Function\[T\]\) \(T, error\)](<#Execute>)
- [type Function](<#Function>)


## Variables

<a name="ErrTimeoutReached"></a>ErrTimeoutReached is returned when the timeout is reached.

```go
var ErrTimeoutReached = errors.New("timeout reached")
```

<a name="Execute"></a>
## func [Execute](<https://github.com/atomicgo/timeout/blob/main/timeout.go#L17>)

```go
func Execute[T any](duration time.Duration, fn Function[T]) (T, error)
```

Execute executes a function and returns the result or an error if the timeout is reached. If the timeout is reached, the Function will be interrupted. If the Function returns an error, the error will be returned.

<a name="Function"></a>
## type [Function](<https://github.com/atomicgo/timeout/blob/main/timeout.go#L9>)

Function is a function that returns a generic value and an error.

```go
type Function[T any] func() (T, error)
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
