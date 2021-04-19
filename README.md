# gofloat

[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/senpathi/gofloat)
[![build](https://github.com/senpathi/gofloat/workflows/build/badge.svg)](https://github.com/senpathi/gofloat/actions)
[![Coverage](https://codecov.io/gh/senpathi/gofloat/branch/master/graph/badge.svg)](https://codecov.io/gh/senpathi/gofloat)
[![Releases](https://img.shields.io/github/release/senpathi/gofloat/all.svg?style=flat-square)](https://github.com/senpathi/gofloat/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/senpathi/gofloat)](https://goreportcard.com/report/github.com/senpathi/gofloat)
[![LICENSE](https://img.shields.io/github/license/senpathi/gofloat.svg?style=flat-square)](https://github.com/senpathi/gofloat/blob/master/LICENSE)

Gofloat is a library that does floating-point number calculations with fixed precision.

## Description
### why goflat
* without gofloat
```go
package main

import "fmt"

func main() {
	f1 := 0.1
	f2 := 0.2
	if f1+f2 == 0.3 {
		fmt.Println(`0.1 + 0.2 = `, f1+f2)
	} else {
		fmt.Println(`0.1 + 0.2 != `, f1+f2)
	}
	//Output: 0.1 + 0.2 != 0.30000000000000004
}
```

* with gofloat
```go
package main

import (
	"fmt"
	"github.com/senpathi/gofloat"
)

func main() {
	f1 := gofloat.ToFloat(0.1, 1)
	f2 := gofloat.ToFloat(0.2, 1)
	if f1.Add(f2).Float64() == 0.3 {
		fmt.Println(`0.1 + 0.2 = `, f1.Add(f2).Float64())
	} else {
		fmt.Println(`0.1 + 0.2 != `, f1.Add(f2).Float64())
	}
	//Output: 0.1 + 0.2 = 0.3
}

```

* gofloat makes sure to do calculations on floating-point numbers with fix number of precision digits
* Converts floating-point numbers into gofloat's `Float` type with sent precision length
* Do Calculations in gofloat's `Float` and then can be converted back into floating-point

## Usage
Documentation is available via
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/senpathi/gofloat).
````go
package main

import (
	"fmt"
	"github.com/senpathi/gofloat"
	"math"
)

func main() {

	// convert Pi and Sqrt(2) into Float with three precisions
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)
	fmt.Println(pi.Float64(), sqrt2.Float64()) //Output: 3.142 1.414

	// add Sqrt(2) to Pi and get results with 3 precisions
	f := pi.Add(sqrt2)
	fmt.Println(f.Float64()) //Output: 4.556

	// subtract Sqrt(2) from Pi and get result with 3 precision
	f = pi.Sub(sqrt2)
	fmt.Println(f.Float64()) //output: 1.728

	// multiply Pi by Sqrt(2) and get result with 3 precision
	f = pi.Multiply(sqrt2)
	fmt.Println(f.Float64()) //Output: 4.443

	// divide Pi by Sqrt(2) and get result with 3 precision
	f = pi.Divide(sqrt2)
	fmt.Println(f.Float64()) //Output: 2.222
}
````
Please see the example programs in the examples directory for reference.

## Limitations
* if the calculated result is larger than maximum float64 value, result will be incorrect.

