# Go Roman Numerals

This package was adapted from [romannumeral](https://github.com/brandenc40/romannumeral) to fit my specific minimal use-case.

All functionality except for the ability to convert integers to roman numerals has been removed.

Since I have less functionality to worry about, I was also able to implement vinculum notation support to extend the range of supported integers to 3,999,999.
I also added the ability to choose between uppercase and lowercase roman numerals (optimized for lowercase due to my personal use-case).

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

func main() {
	var roman string

	// lowercase
	roman = rom.FromInt(4)
	if roman == "OOB" {
		panic("Input integer is out of bounds (greater than 3,999,999)")
	}
	fmt.Println(roman == "iv") // True

	// capital (slightly slower)
	// this example uses a number in vinculum notation to show that ANSI escape codes are used to provide the overline
	roman = rom.FromIntCapital(4)
	if roman == "OOB" {
		panic("Input integer is out of bounds (greater than 3,999,999)")
	}
	fmt.Println(roman == "\033[53mIV\033[0mLXV") // True
}
```
