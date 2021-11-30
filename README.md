# go-dirsize

Provides a function to calculate the directory size.

## Installation

```
go get -u github.com/kenkyu392/go-dirsize
```

## Usage

```go
package main

import (
	"log"

	"github.com/kenkyu392/go-dirsize"
)

func main() {
	size, err := dirsize.Calc(
		// Path of the folder to calculate.
		"./testdata",
		// Byte unit of the calculation result (B/KB/MB/GB).
		dirsize.MB,
		// Allows you to set exclusion rule options.
		// The following implementation excludes files and
		// folders with separate rules.
		// You can also write: dirsize.Ignore("ignore*")
		dirsize.IgnoreFile("ignore_*"),
		dirsize.IgnoreDir("ignore"),
	)
	log.Println(size, err)
}
```

## License

[MIT](LICENSE)
