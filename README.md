
# stopwatch

A simply package for timing your code. The intention is to provide a simple,
light-weight library for benchmarking specific bits of your code when need be.

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/rodevl/stopwatch"
)

func main() {
	watch := stopwatch.Start()

	for i := 0; i < 10; i++ {
		// Do some work.
		time.Sleep(time.Second)
		fmt.Printf("Current stopwatch value in seconds: %d\n", 
			watch.CurrentDurationSecond())
	}

	watch.Stop()

	fmt.Printf("Elapsed stopwatch value in seconds: %d\n", watch.Seconds())
}
```

## Contributing

1. Fork and fix/implement in a branch.
1. Make sure tests pass.
1. Make sure you've added new coverage.
1. Submit a PR.
