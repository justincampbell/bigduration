# bigduration [![Build Status](https://travis-ci.org/justincampbell/bigduration.svg?branch=master)](https://travis-ci.org/justincampbell/bigduration) [![GoDoc](https://godoc.org/github.com/justincampbell/bigduration?status.svg)](https://godoc.org/github.com/justincampbell/bigduration)

Additional [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration) suffixes for **D**ays, **W**eeks, **M**onths, and **Y**ears

## Installation

```
go get github.com/justincampbell/bigduration
```

## Example

```go
package main

import (
	"fmt"

	"github.com/justincampbell/bigduration"
)

func main() {
	examples := []string{
		"1D",
		"1W",
		"1M",
		"1Y",
		"4Y3M2D1h30m15s",
	}

	for _, example := range examples {
		duration, err := bigduration.ParseDuration(example)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s -> %s\n", example, duration)
	}
}
```

```
$ go run main.go
1D -> 24h0m0s
1W -> 168h0m0s
1M -> 720h0m0s
1Y -> 8760h0m0s
4Y3M2D1h30m15s -> 37249h30m15s
```
