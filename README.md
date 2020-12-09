# etag

[![GitHub release](https://img.shields.io/github/release/gaius-qi/etag.svg)](https://github.com/gaius-qi/etag/releases)
[![Build Status](https://travis-ci.org/gaius-qi/etag.svg?branch=master)](https://travis-ci.org/gaius-qi/etag)
[![Coverage Status](https://coveralls.io/repos/github/gaius-qi/etag/badge.svg?branch=master)](https://coveralls.io/github/gaius-qi/etag?branch=master)
[![GoDoc](https://godoc.org/github.com/gaius-qi/etag?status.svg)](https://godoc.org/github.com/gaius-qi/etag)

Generate a simple ETag

## Installation

```shell
$ go get github.com/gaius-qi/etag
```

## Usage

```js
package main

import (
    "fmt"

    "github.com/gaius-qi/etag"
)

func main() {
    entity := "Hello World!"
    fmt.Println("Etag:", etag.Generate(entity, false))
    fmt.Println("Week Etag:", etag.Generate(entity, true))
}
```

## Issues

- [Open an issue in GitHub](https://github.com/gaius-qi/etag/issues)

## License

[MIT](LICENSE)
