package gofloat_test

import (
	"fmt"
	"math"

	"github.com/senpathi/gofloat"
)

func ExampleToFloat() {
	// convert Pi and Sqrt(2) into Float with three precisions
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	fmt.Println(pi.Float64(), sqrt2.Float64())
	//Output: 3.142 1.414
}

func ExampleFloat_Add() {
	// add Sqrt(2) to Pi and get results with 3 precisions
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	f := pi.Add(sqrt2)
	fmt.Println(`PI + Sqrt(2) =`, f.Float64())
	//Output: PI + Sqrt(2) = 4.556
}

func ExampleFloat_Sub() {
	// subtract Sqrt(2) from Pi and get result with 3 precision
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	f := pi.Sub(sqrt2)
	fmt.Println(`PI - Sqrt(2) =`, f.Float64())
	//Output: PI - Sqrt(2) = 1.728
}

func ExampleFloat_Multiply() {
	// multiply Pi by Sqrt(2) and get result with 3 precision
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	f := pi.Multiply(sqrt2)
	fmt.Println(`PI * Sqrt(2) =`, f.Float64())
	//Output: PI * Sqrt(2) = 4.443
}

func ExampleFloat_Divide() {
	// divide Pi by Sqrt(2) and get result with 3 precision
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	f := pi.Divide(sqrt2)
	fmt.Println(`PI / Sqrt(2) =`, f.Float64())
	//Output: PI / Sqrt(2) = 2.222
}

func ExampleFloat_Float64() {
	// convert Pi and Sqrt(2) into Float with three precisions and convert back to float64
	pi := gofloat.ToFloat(math.Pi, 3)
	sqrt2 := gofloat.ToFloat(math.Sqrt2, 3)

	fmt.Println(pi.Float64(), sqrt2.Float64())
	//Output: 3.142 1.414
}
