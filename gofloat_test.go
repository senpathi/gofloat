package domain

import (
	"fmt"
	"math"
	"testing"
)

func TestCurrency(t *testing.T) {
	t.Run(`TestCurrencyAdd`, func(t *testing.T) {
		f1 := 1.3725
		f2 := 2.4725
		x := ToCurrency(f1)
		y := ToCurrency(f2)
		x = x.Add(y)
		if x.Float64() != math.Round((f1+f2)*float64(floatLength))/float64(floatLength) {
			t.Error(fmt.Sprintf(`expected %f, received %f`, f1+f2, x.Float64()))
		}
	})

	t.Run(`TestCurrencySub`, func(t *testing.T) {
		f1 := -1.3725
		f2 := 2.47255
		x := ToCurrency(f1)
		y := ToCurrency(f2)
		x = x.Sub(y)
		expected := math.Round((f1-f2)*float64(floatLength)) / float64(floatLength)
		if x.Float64() != expected {
			t.Error(fmt.Sprintf(`expected %f, received %f`, expected, x.Float64()))
		}
	})

	t.Run(`TestCurrencyMultiply`, func(t *testing.T) {
		f1 := 1.3725
		f2 := 2.4725
		x := ToCurrency(f1)
		y := ToCurrency(f2)
		x = x.Multiply(y)
		if x.Float64() != math.Round(f1*f2*float64(floatLength))/float64(floatLength) {
			t.Error(fmt.Sprintf(`expected %f, received %f`, f1*f2, x.Float64()))
		}
	})

	t.Run(`TestCurrencyDivide`, func(t *testing.T) {
		f1 := 1.3725
		f2 := 2.4725
		x := ToCurrency(f1)
		y := ToCurrency(f2)
		x = x.Divide(y)
		if x.Float64() != math.Round(f1/f2*float64(floatLength))/float64(floatLength) {
			t.Error(fmt.Sprintf(`expected %f, received %f`, f1/f2, x.Float64()))
		}
	})
}
