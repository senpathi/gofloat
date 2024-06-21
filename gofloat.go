package gofloat

import (
	"math"
	"strconv"
)

// Float is to keep float's integer value, decimal value and its precision
type Float struct {
	integer         int64
	decimal         int64
	precision       int
	precisionFactor int64
}

// ToFloat returns a Float of f to given precision
func ToFloat(f float64, precision int) Float {
	precision = int(math.Abs(float64(precision)))
	floatingLength := math.Pow10(precision)

	return toFloat(f, precision, int64(floatingLength))
}

func toFloat(f float64, precision int, precisionFactor int64) Float {
	rounded := math.Round(f * float64(precisionFactor))
	return Float{
		integer:         int64(rounded) / precisionFactor,
		decimal:         int64(rounded) % precisionFactor,
		precision:       precision,
		precisionFactor: precisionFactor,
	}
}

// Add returns a Float of f + x
func (f Float) Add(x Float) Float {
	maxP, maxFact := getMaxPrecision(f, x)
	sum := f.Float64() + x.Float64()
	return toFloat(sum, maxP, maxFact)
}

// Sub returns a Float of f - x
func (f Float) Sub(x Float) Float {
	maxP, maxFact := getMaxPrecision(f, x)
	sub := f.Float64() - x.Float64()
	return toFloat(sub, maxP, maxFact)
}

// Multiply returns a Float of f * x
func (f Float) Multiply(x Float) Float {
	maxP, maxFact := getMaxPrecision(f, x)
	mul := f.Float64() * x.Float64()
	return toFloat(mul, maxP, maxFact)
}

// Divide returns a Float of f / x
//
// Special cases are :
//
//	x.Float64 = 0 returns Float value 0
func (f Float) Divide(x Float) Float {
	if x.Float64() == 0 {
		return Float{
			integer:         0,
			decimal:         0,
			precision:       f.precision,
			precisionFactor: f.precisionFactor,
		}
	}
	maxP, maxFact := getMaxPrecision(f, x)
	dev := f.Float64() / x.Float64()
	return toFloat(dev, maxP, maxFact)
}

// Float64 returns a float64 value of f
func (f Float) Float64() float64 {
	return (float64(f.integer*f.precisionFactor) + float64(f.decimal)) / float64(f.precisionFactor)
}

func (f Float) String() string {
	return strconv.FormatFloat(f.Float64(), 'f', f.precision, 64)
}

func getMaxPrecision(p1, p2 Float) (max int, maxFactor int64) {
	if p1.precision < p2.precision {
		return p2.precision, p2.precisionFactor
	}

	return p1.precision, p1.precisionFactor
}
