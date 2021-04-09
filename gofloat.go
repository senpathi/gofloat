package domain

import (
	"math"
)

type Currency struct {
	integer int64
	decimal int64
}

var floatLength int64 = 1e4

func DefineFloatingPointLength(length int) {
	floatLength = int64(math.Pow10(length))
}

func ToCurrency(f float64) Currency {
	rounded := math.Round(f * float64(floatLength))
	return Currency{
		integer: int64(rounded) / floatLength,
		decimal: int64(rounded) % floatLength,
	}
}

func (c Currency) Add(x Currency) Currency {
	d := c.decimal + x.decimal

	return Currency{
		integer: d/floatLength + c.integer + x.integer,
		decimal: d % floatLength,
	}
}

func (c Currency) Sub(x Currency) Currency {
	ci := c.integer
	cd := c.decimal
	xi := x.integer
	xd := x.decimal

	if ci*floatLength+cd < xi*floatLength+xd {
		if cd >= 0 && xd >= 0 && cd > xd {
			xd = xd + floatLength
			xi -= 1
		}
		d := xd - cd
		return Currency{
			integer: -(d/floatLength + xi - ci),
			decimal: -(d % floatLength),
		}
	}

	if cd >= 0 && xd >= 0 && xd > cd {
		cd = cd + floatLength
		ci -= 1
	}
	d := cd - xd
	return Currency{
		integer: d/floatLength + ci - xi,
		decimal: d % floatLength,
	}
}

func (c Currency) Multiply(x Currency) Currency {
	d := (c.decimal + (c.integer * floatLength)) * (x.decimal + (x.integer * floatLength))

	return Currency{
		integer: d / 1e8,
		decimal: (d % 1e8) / floatLength,
	}
}

func (c Currency) Divide(x Currency) Currency {
	c1 := c.integer*floatLength + c.decimal
	x1 := x.integer*floatLength + x.decimal
	decimal := int64(0)
	d := c1 % x1
	for i := floatLength / 10; i >= 1; i = i / 10 {
		decimal += (d * 10 / x1) * i
		d = d * 10 % x1
	}

	return Currency{
		integer: c1 / x1,
		decimal: decimal,
	}
}

func (c Currency) Float64() float64 {
	return (float64(c.integer*floatLength) + float64(c.decimal)) / float64(floatLength)
}
