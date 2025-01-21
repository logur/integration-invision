> [!WARNING]
> This project is archived and no longer maintained. Consider using [`log/slog`](https://pkg.go.dev/log/slog) instead.
>
> Read more about why it was archived in [this post](https://sagikazarmark.com/blog/posts/less-is-more-archive-projects-for-a-better-open-source-ecosystem/).

# Logur integration for [InVision go-logger](https://github.com/InVisionApp/go-logger) interface

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/integration-invision/CI?style=flat-square)](https://github.com/logur/integration-invision/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/logur/integration-invision?style=flat-square)](https://codecov.io/gh/logur/integration-invision)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/invision?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/invision)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/logur.dev/integration/invision)


## Installation

```bash
go get logur.dev/integration/invision
```


## Usage

```go
package main

import (
	"logur.dev/logur"
	invisionintegration "logur.dev/integration/invision"
)

func main() {
	logger := invisionintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
