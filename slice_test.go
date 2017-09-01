package gosc

import (
	"testing"
)

// TestEqStringSlices tests the EqSlices function with strings slice
func TestEqStringSlice(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		a        []string
		b        []string
		expected bool
	}{
		{[]string{"foo", "bar"}, []string{"bar", "foo"}, false},
		{[]string{"foo", "bar"}, []string{"bar"}, false},
		{[]string{"foo", "bar"}, []string{"foo"}, false},
		{[]string{"foo", "bar"}, []string{"foo", "bar"}, true},
		{[]string{"foo", "bar"}, []string{"\u0066\u006f\u006f", "\u0062\u0061\u0072"}, true},
		{[]string{"foo", "bar"}, []string{"\x66\x6f\x6f", "\x62\x61\x72"}, true},
		{[]string{"foo", "bar"}, []string{"\u0066\u006f\u006f", "\x62\x61\x72"}, true},
		{[]string{"foo", "bar"}, []string{"\x66\x6f\x6f", "bar"}, true},
	}

	for _, test := range tests {
		actual := EqSlices(&test.a, &test.b)
		if actual != test.expected {
			t.Errorf("Expected EqSlices(%q, %q) to be %v, got %v", test.a, test.a, test.expected, actual)
		}
	}
}

// TestEqIntSlices tests the EqSlices function with strings slice
func TestEqIntSlice(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		a        []int
		b        []int
		expected bool
	}{
		{[]int{5, -3}, []int{-3, 5}, false},
		{[]int{5, -3}, []int{-5, 3}, false},
		{[]int{5, -3}, []int{-5, -3}, false},
		{[]int{5, -3}, []int{5}, false},
		{[]int{5, -3}, []int{5, -3, 1}, false},
		{[]int{5, -3}, []int{5, -3}, true},
	}

	for _, test := range tests {
		actual := EqSlices(&test.a, &test.b)
		if actual != test.expected {
			t.Errorf("Expected EqSlices(%q, %q) to be %v, got %v", test.a, test.a, test.expected, actual)
		}
	}
}

// TestInStringSlice tests the InSlice function with strings slice
func TestInStringSlice(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []string
		needle   string
		expected bool
	}{
		{[]string{"foo", "bar"}, "bar", true},
		{[]string{"foo", "bar"}, "", false},
		{[]string{"foo", "bar"}, "0", false},
		{[]string{"foo", "bar"}, "b", false},
		{[]string{"a", "b"}, "\u0061", true},
	}

	for _, test := range tests {
		actual := InSlice(test.needle, &test.haystack)
		if actual != test.expected {
			t.Errorf("Expected InSlice(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}

// TestInIntSlice tests the InSlice function with ints slice
func TestInIntSlice(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []int
		needle   int
		expected bool
	}{
		{[]int{0, 1, 2}, 2, true},
		{[]int{0, 1, 2}, 0, true},
		{[]int{0, 1, 2, -5}, -5, true},
		{[]int{0, 1, 2, -5}, 5, false},
	}

	for _, test := range tests {
		actual := InSlice(test.needle, &test.haystack)
		if actual != test.expected {
			t.Errorf("Expected InSlice(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}
