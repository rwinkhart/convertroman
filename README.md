# Go Roman Numerals

This package was adapted from [romannumeral](https://github.com/brandenc40/romannumeral) to fit my specific minimal use-case.

All functionality except for the ability to convert integers to **lowercase** roman numerals has been removed.

If you need any additional functionality or further documentation, please see the original package.

### Benchmark Results

```sh
goos: linux
goarch: amd64
pkg: github.com/rwinkhart/convertroman
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkFromInt-24    	61993154	        20.82 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rwinkhart/convertroman	1.827s
```

### Example

```go
package main

import (
	"fmt"
	rom "github.com/rwinkhart/convertroman"
)

func ExampleFromInt() {
	roman := FromInt(4)
	if roman == "OOB" {
		panic("Input integer is out of bounds")
	}
	fmt.Println(roman == "IV") // True
}
```
