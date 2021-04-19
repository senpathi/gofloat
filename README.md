# gofloat
Gofloat is a library that does floating-point number calculations with fixed precision.

## Description
* Converts floating-point numbers to gofloat type with sent precision length
* Do Calculations in converted type and it can be converted back to floating-point

## Usage
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

