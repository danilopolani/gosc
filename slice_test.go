package gosc

import (
	"strings"
	"testing"
)

// TestAnyString tests the AnyString function
func TestAnyString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		s        []string
		expected bool
	}{
		{[]string{"foo", "bar", "baz"}, true},
		{[]string{"foo", "\u0062\u0061\u0072", "baz"}, true},
		{[]string{"foo", "bar", "buz"}, true},
		{[]string{"foo", "bur", "buz"}, false},
	}

	for _, test := range tests {
		actual := AnyString(test.s, func(s string) bool {
			return strings.HasPrefix(s, "ba")
		})
		if actual != test.expected {
			t.Errorf("Expected AnyString(%q, fn) to be %v, got %v", test.s, test.expected, actual)
		}
	}
}

// TestAnyInt tests the AnyInt function
func TestAnyInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		s        []int
		expected bool
	}{
		{[]int{0, 2, 4}, true},
		{[]int{}, false},
		{[]int{2, 4, 5}, true},
		{[]int{5}, false},
		{[]int{5, 2}, true},
		{[]int{-2, 4}, true},
	}

	for _, test := range tests {
		actual := AnyInt(test.s, func(i int) bool {
			return i%2 == 0
		})
		if actual != test.expected {
			t.Errorf("Expected AnyInt(%q, fn) to be %v, got %v", test.s, test.expected, actual)
		}
	}
}

// TestIndexi tests the Indexi function
func TestIndexi(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []string
		needle   string
		expected int
	}{
		{[]string{"FOO", "bAr"}, "bar", 1},
		{[]string{"FOO", "bAr"}, "foo", 0},
		{[]string{"FOO", "bAr"}, "nil", -1},
		{[]string{"FOO", "bAr"}, "\u0062\u0061\u0072", 1},
		{[]string{"FOO", "bAr"}, "0", -1},
		{[]string{"FOO", "bAr"}, "", -1},
	}

	for _, test := range tests {
		actual := Indexi(test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("Expected Indexi(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}

// TestStringIndex tests the Index function with strings slice
func TestStringIndex(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []string
		needle   string
		expected int
	}{
		{[]string{"foo", "bar"}, "bar", 1},
		{[]string{"foo", "bar"}, "foo", 0},
		{[]string{"foo", "bar"}, "nil", -1},
		{[]string{"foo", "bar"}, "\u0062\u0061\u0072", 1},
		{[]string{"foo", "bar"}, "0", -1},
		{[]string{"foo", "bar"}, "", -1},
	}

	for _, test := range tests {
		actual := Index(&test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("Expected Index(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}

// TestIntIndex tests the Index function with ints slice
func TestIntIndex(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []int
		needle   int
		expected int
	}{
		{[]int{1, -5}, -5, 1},
		{[]int{1, -5}, 0, -1},
		{[]int{1, -5}, 1, 0},
	}

	for _, test := range tests {
		actual := Index(&test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("Expected Index(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}

// TestFloatIndex tests the Index function with ints slice
func TestFloatIndex(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		haystack []float64
		needle   float64
		expected int
	}{
		{[]float64{1, 6.3, -5.7}, -5.7, 2},
		{[]float64{1, 6.3, -5.7}, 0, -1},
		{[]float64{1, 6.3, -5.7}, 1, 0},
	}

	for _, test := range tests {
		actual := Index(&test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("Expected Index(%q, %q) to be %v, got %v", test.haystack, test.needle, test.expected, actual)
		}
	}
}

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
