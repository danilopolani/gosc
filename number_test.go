package gosc

import (
	"testing"
)

// TestIsInt tests the IsInt function
func TestIsInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주", false},
		{"nil", false},
		{"null", false},
		{"1", true},
		{"0", true},
		{"1111-", false},
		{"-1", true},
		{"5.3", false},
		{"\u0031", true},
		{"\x31", true},
	}

	for _, test := range tests {
		actual := IsInt(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsInt(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsFloat tests the IsFloat function
func TestIsFloat(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주", false},
		{"nil", false},
		{"null", false},
		{"1", true},
		{"0", true},
		{"1111-", false},
		{"-1", true},
		{"-1.65", true},
		{"5.3", true},
		{"\u0031", true},
		{"\x31", true},
		{"\u0035\u002e\u0033", true},
		{"\x35\x2e\x33", true},
	}

	for _, test := range tests {
		actual := IsFloat(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsFloat(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}
