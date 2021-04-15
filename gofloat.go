package domain

import (
	"math"
)

type Float struct {
	integer   int64
	decimal   int64
	precision int
}

// ToFloat returns a Float of f to given precision
func ToFloat(f float64, precision int) Float {
	floatingLength := int64(math.Trunc(math.Pow(10, float64(precision))))
	rounded := math.Round(f * float64(floatingLength))

	return Float{
		integer:   int64(rounded) / floatingLength,
		decimal:   int64(rounded) % floatingLength,
		precision: precision,
	}
}

// Add returns a Float of f + x
func (f Float) Add(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	sum := f.Float64() + x.Float64()
	return ToFloat(sum, max)
}

// Sub returns a Float of f - x
func (f Float) Sub(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	sub := f.Float64() - x.Float64()
	return ToFloat(sub, max)
}

// Multiply returns a Float of f * x
func (f Float) Multiply(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	mul := f.Float64() * x.Float64()
	return ToFloat(mul, max)
}

// Divide returns a Float of f / x
//
// Special cases are :
//  x.Float64 = 0 returns Float value 0
func (f Float) Divide(x Float) Float {
	if x.Float64() == 0 {
		return ToFloat(0, 0)
	}
	max := getMaxPrecision(f.precision, x.precision)
	dev := f.Float64() / x.Float64()
	return ToFloat(dev, max)
}

// Float64 returns a float64 value of f
func (f Float) Float64() float64 {
	floatingLength := int64(math.Trunc(math.Pow(10, float64(f.precision))))
	return (float64(f.integer*floatingLength) + float64(f.decimal)) / float64(floatingLength)
}

func getMaxPrecision(p1, p2 int) (max int) {
	max = p1
	if p1 < p2 {
		max = p2
	}
	return
}
