package main

import (
	"fmt"
	"calc/calc"
	"calc/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	expression, simple string
}

func TestCalculator(t *testing.T) {
	tests := []testcase{
		{
			expression: "2 * 3 + 1",
			simple:     "7",
		},
		{
			expression: "1/2 * 3_3/4",
			simple:     "1_7/8",
		},
		{
			expression: "2_3/8 + 9/8",
			simple:     "3_1/2",
		},
		{
			expression: "3/4 + 1 * 3_1/2",
			simple:     "4_1/4",
		},
		{
			expression: "3/4 - 4 / 3_1/4 * 3",
			simple:     "-153/52",
		},
		{
			expression: "1 + 2",
			simple:     "3",
		},
		{
			expression: "1 + 2 * 3",
			simple:     "7",
		},
		{
			expression: " 2 + 2 + 2 + 3",
			simple:     "9",
		},
		{
			expression: " 2 + 2 + 2 + 6 / 2 ",
			simple:     "9",
		},
		{
			expression: " 2 + 4 / 2 + 2 + 3 ",
			simple:     "9",
		},
	}

	for idx, test := range tests {
		var calculator calc.CalcStack

		parser.Parse(&calculator, test.expression)

		res, err := calculator.Pop()
		if err != nil {
			panic(err)
		}

		assert.Equal(t, test.simple, res.Simplify(), fmt.Sprintf("Simplified fraction should be equal to %s for test case index %d", test.simple, idx))
	}
}
