package domain

import (
	"math"
)

type Float struct {
	integer   int64
	decimal   int64
	precision int
}

func ToFloat(f float64, precision int) Float {
	floatingLength := int64(math.Trunc(math.Pow(10, float64(precision))))
	rounded := math.Round(f * float64(floatingLength))

	return Float{
		integer:   int64(rounded) / floatingLength,
		decimal:   int64(rounded) % floatingLength,
		precision: precision,
	}
}

func (f Float) Add(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	sum := f.Float64() + x.Float64()
	return ToFloat(sum, max)
}

func (f Float) Sub(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	sub := f.Float64() - x.Float64()
	return ToFloat(sub, max)
}

func (f Float) Multiply(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	mul := f.Float64() * x.Float64()
	return ToFloat(mul, max)
}

func (f Float) Divide(x Float) Float {
	max := getMaxPrecision(f.precision, x.precision)
	dev := f.Float64() / x.Float64()
	return ToFloat(dev, max)
}

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
