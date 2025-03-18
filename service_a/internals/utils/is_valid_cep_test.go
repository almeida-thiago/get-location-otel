package utils

import (
	"testing"
)

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		cep      string
		expected bool
	}{
		{"12345678", true},
		{"87654321", true},
		{"1234", false},
		{"abcdefgh", false},
		{"1234-5678", false},
		{"123456789", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.cep, func(t *testing.T) {
			result := IsValidCEP(test.cep)
			if result != test.expected {
				t.Errorf("IsValidCEP(%s) = %v; want %v", test.cep, result, test.expected)
			}
		})
	}
}
