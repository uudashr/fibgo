[![Build Status](https://travis-ci.org/uudashr/fibgo.svg?branch=master)](https://travis-ci.org/uudashr/fibgo)
[![Coverage Status](https://coveralls.io/repos/github/uudashr/fibgo/badge.svg?branch=master)](https://coveralls.io/github/uudashr/fibgo?branch=master)

# Fibonacci

Provide function for fibonacci operation

## Installation
```shell
$ go get github.com/uudashr/fibgo
```

## Usage
Get fibonacci number of N
```go
import (
  "fmt"

  fib "github.com/uudashr/fibgo"
)

func main() {
  fmt.Println(fib.N(10))
}
```

Will result
```
7
```

The N will start from 0.

Or get sequence with length 10

```go
import (
  "fmt"

  fib "github.com/uudashr/fibgo"
)

func main() {
  fmt.Println(fib.Seq(10))
}
```

Will result
```
[0 1 1 2 3 5 8 13 21 34]
```
