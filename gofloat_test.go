package gofloat

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestFloat(t *testing.T) {
	t.Run(`TestToFloat`, func(t *testing.T) {
		testTable := setupTestTable(`NONE`)
		for _, tt := range testTable {
			x := ToFloat(tt.input1, tt.precision1)
			if x.Float64() != tt.out {
				t.Error(fmt.Sprintf(`expected %f, received %f`, tt.out, x.Float64()))
			}
		}
	})

	t.Run(`TestFloatAdd`, func(t *testing.T) {
		testTable := setupTestTable(`ADD`)
		for _, tt := range testTable {
			x := ToFloat(tt.input1, tt.precision1)
			y := ToFloat(tt.input2, tt.precision2)
			x = x.Add(y)
			if x.Float64() != tt.out {
				t.Error(fmt.Sprintf(`expected %f + %f = %f, but received %f with precisions %d and %d`,
					tt.input1, tt.input2, tt.out, x.Float64(), tt.precision1, tt.precision2))
			}
		}
	})

	t.Run(`TestFloatSub`, func(t *testing.T) {
		testTable := setupTestTable(`SUB`)
		for _, tt := range testTable {
			x := ToFloat(tt.input1, tt.precision1)
			y := ToFloat(tt.input2, tt.precision2)
			x = x.Sub(y)
			if x.Float64() != tt.out {
				t.Error(fmt.Sprintf(`expected %f - %f = %f, but received %f with precision %d and %d`,
					tt.input1, tt.input2, tt.out, x.Float64(), tt.precision1, tt.precision2))
			}
		}
	})

	t.Run(`TestFloatMultiply`, func(t *testing.T) {
		testTable := setupTestTable(`MUL`)
		for _, tt := range testTable {
			x := ToFloat(tt.input1, tt.precision1)
			y := ToFloat(tt.input2, tt.precision2)
			x = x.Multiply(y)
			if x.Float64() != tt.out {
				t.Error(fmt.Sprintf(`expected %f * %f = %f, but received %f with precision %d and %d`,
					tt.input1, tt.input2, tt.out, x.Float64(), tt.precision1, tt.precision2))
			}
		}
	})

	t.Run(`TestFloatDivide`, func(t *testing.T) {
		testTable := setupTestTable(`DEV`)
		for _, tt := range testTable {
			x := ToFloat(tt.input1, tt.precision1)
			y := ToFloat(tt.input2, tt.precision2)
			x = x.Divide(y)
			if x.Float64() != tt.out {
				t.Error(fmt.Sprintf(`expected %f / %f = %f, but received %f with precision %d and %d`,
					tt.input1, tt.input2, tt.out, x.Float64(), tt.precision1, tt.precision2))
			}
		}
	})

	t.Run(`TestFloatDivideByZero`, func(t *testing.T) {
		x := ToFloat(100.1234, 2)
		y := ToFloat(0, 0)
		x = x.Divide(y)
		if x.Float64() != 0 {
			t.Error(fmt.Sprintf(`expected %f / %f = %f, but received %f`,
				100.1234, 0.0, 0.0, x.Float64()))
		}
	})

	t.Run(`TestFloatString`, func(t *testing.T) {
		type stringTest struct {
			input     float64
			precision int
			output    string
		}

		testTable := []stringTest{
			{1.234567, 6, `1.234567`},
			{-1.234567, 6, `-1.234567`},
		}
		for _, tt := range testTable {
			x := ToFloat(tt.input, tt.precision)
			if x.String() != tt.output {
				t.Error(fmt.Sprintf(`expected %s, received %s`, tt.output, x.String()))
			}
		}
	})

}

func BenchmarkFloat(b *testing.B) {
	b.Run(`Bench_ToFloat`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToFloat(float64(i/33), 5)
		}
	})
	b.Run(`Bench_Float_Add`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToFloat(float64(i/33), 5).Add(ToFloat(float64(i/66), 5))
		}
	})
	b.Run(`Bench_Float_Sub`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToFloat(float64(i/33), 5).Sub(ToFloat(float64(i/66), 5))
		}
	})
	b.Run(`Bench_Float_Multiply`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToFloat(float64(i/33), 5).Multiply(ToFloat(float64(i/66), 5))
		}
	})
	b.Run(`Bench_Float_Divide`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToFloat(float64(i/33), 5).Divide(ToFloat(float64(i/66), 5))
		}
	})
}

type TestInput struct {
	input1     float64
	input2     float64
	precision1 int
	precision2 int
	out        float64
}

func setupTestTable(operator string) []TestInput {
	inputs := make([]TestInput, 1000)

	for i := range inputs {
		inputs[i] = TestInput{
			input1:     rand.Float64() * 1e5 * math.Pow(-1, float64(i/250)),
			input2:     rand.Float64() * 1e5 * math.Pow(-1, float64(i/500)),
			precision1: rand.Intn(9),
			precision2: rand.Intn(9),
		}

		length1 := math.Trunc(math.Pow10(inputs[i].precision1))

		f1 := ToFloat(inputs[i].input1, inputs[i].precision1)
		rounded1 := f1.Float64()
		f2 := ToFloat(inputs[i].input2, inputs[i].precision2)
		rounded2 := f2.Float64()
		_, maxFact := getMaxPrecision(f1, f2)
		lengthMax := float64(maxFact)

		switch operator {
		case `NONE`:
			inputs[i].out = math.Round(inputs[i].input1*length1) / length1
		case `ADD`:
			inputs[i].out = math.Round((rounded1+rounded2)*lengthMax) / lengthMax
		case `SUB`:
			inputs[i].out = math.Round((rounded1-rounded2)*lengthMax) / lengthMax
		case `MUL`:
			inputs[i].out = math.Round((rounded1*rounded2)*lengthMax) / lengthMax
		case `DEV`:
			inputs[i].out = math.Round((rounded1/rounded2)*lengthMax) / lengthMax
		}
	}
	return inputs
}
