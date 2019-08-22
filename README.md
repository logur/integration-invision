# Logur integration for [InVision go-logger](https://github.com/InVisionApp/go-logger) interface

[![CircleCI](https://circleci.com/gh/logur/integration-invision.svg?style=svg)](https://circleci.com/gh/logur/integration-invision)
[![Coverage](https://gocover.io/_badge/logur.dev/integration/invision)](https://gocover.io/logur.dev/integration/invision)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/invision?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/invision)
[![GolangCI](https://golangci.com/badges/github.com/logur/integration-invision.svg)](https://golangci.com/r/github.com/logur/integration-invision)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/integration-invision)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/integration/invision)


## Installation

```bash
go get logur.dev/integration/invision
```


## Usage

```go
package main

import (
	"github.com/goph/logur"
	invisionintegration "logur.dev/integration/invision"
)

func main() {
	logger := invisionintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
