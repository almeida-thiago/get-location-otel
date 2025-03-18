package utils

import (
	"math"
	"testing"
)

func TestConvertCToF(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.6},
		{25, 77},
	}

	for _, test := range tests {
		t.Run("ConvertCToF", func(t *testing.T) {
			result := ConvertCToF(test.celsius)
			if math.Abs(result-test.expected) > 0.01 {
				t.Errorf("ConvertCToF(%v) = %v; want %v", test.celsius, result, test.expected)
			}
		})
	}
}

func TestConvertCToK(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 273},
		{100, 373},
		{-40, 233},
		{37, 310},
		{25, 298},
	}

	for _, test := range tests {
		t.Run("ConvertCToK", func(t *testing.T) {
			result := ConvertCToK(test.celsius)
			if math.Abs(result-test.expected) > 0.01 {
				t.Errorf("ConvertCToK(%v) = %v; want %v", test.celsius, result, test.expected)
			}
		})
	}
}
