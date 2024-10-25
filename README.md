# Go Roman Numerals

This package was adapted from [romannumeral](https://github.com/brandenc40/romannumeral) to fit my specific minimal use-case.

All functionality except for the ability to convert integers to roman numerals has been removed.

If you need any additional functionality or further documentation, please see the original package.

### Benchmark Results

```sh
goos: linux
goarch: amd64
pkg: github.com/rwinkhart/convertroman
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkFromInt-24    	62046414	        21.08 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rwinkhart/convertroman	1.332s
```

### Example

```go
package main

import (
	"fmt"
	rom "github.com/rwinkhart/convertroman"
)

func ExampleFromInt() {
	roman, err := rom.FromInt(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman == "IV") // True
}
```
