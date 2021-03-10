# etag

[![GitHub release](https://img.shields.io/github/release/gaius-qi/etag.svg)](https://github.com/gaius-qi/etag/releases)
[![Github Build Status](https://github.com/gaius-qi/etag/workflows/Go/badge.svg?branch=main)](https://github.com/gaius-qi/etag/actions?query=workflow%3AGo+branch%3Amain)
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
