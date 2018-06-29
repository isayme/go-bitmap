## Bitmap
A bitmap implement in Go.

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