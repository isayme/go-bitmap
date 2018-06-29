## Bitmap
A bitmap implement in Go.
[![Build Status](https://travis-ci.org/isayme/go-bitmap.svg?branch=master)](https://travis-ci.org/isayme/go-bitmap)
[![Coverage Status](https://coveralls.io/repos/github/isayme/go-bitmap/badge.svg?branch=master)](https://coveralls.io/github/isayme/go-bitmap?branch=master)

## Example
```
package main

import (
	"fmt"

	bitmap "github.com/isayme/go-bitmap"
)

func main() {
	bitmap := bitmap.New()

	bitmap.Set(1)

	fmt.Println(bitmap.Has(1))
	fmt.Println(bitmap.Has(2))
}
```