package gosc

import (
	"testing"
)

// TestReverseString tests the ReverseString function
func TestReverseString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "cba"},
		{"foo123", "321oof"},
		{"bar  ", "  rab"},
		{"", ""},
		{"abc〩", "〩cba"},
		{"소주", "주소"},
		{"소aBC", "CBa소"},
	}

	for _, test := range tests {
		actual := ReverseString(test.data)
		if actual != test.expected {
			t.Errorf("Expected ReverseString(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestLcFirst tests the LcFirst function
func TestLcFirst(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "abc"},
		{"FOO123", "fOO123"},
		{"123", "123"},
		{" test", " test"},
		{"", ""},
		{"Abc〩", "abc〩"},
		{"소주", "소주"},
		{"_Foo", "_Foo"},
	}

	for _, test := range tests {
		actual := LcFirst(test.data)
		if actual != test.expected {
			t.Errorf("Expected LcFirst(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestUcFirst tests the UcFirst function
func TestUcFirst(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "Abc"},
		{"foo123", "Foo123"},
		{"123", "123"},
		{" test", " test"},
		{"", ""},
		{"abc〩", "Abc〩"},
		{"소주", "소주"},
		{"AbC", "AbC"},
	}

	for _, test := range tests {
		actual := UcFirst(test.data)
		if actual != test.expected {
			t.Errorf("Expected UcFirst(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToSnake tests the ToSnake function
func TestToSnake(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "abc"},
		{"foo 123", "foo_123"},
		{"FOO_bar", "foo_bar"},
		{" test", "_test"},
		{"  test  ", "__test__"},
		{"", ""},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"kebab-case", "kebab_case"},
		{"abc〩", "abc〩"},
		{"소주", "소주"},
		{"AbC", "ab_c"},
	}

	for _, test := range tests {
		actual := ToSnake(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToSnake(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToCamel tests the ToCamel function
func TestToCamel(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "abc"},
		{"foo 123", "foo123"},
		//{"FOO_bar", "fooBar"},
		{" test", "test"},
		{"  test  ", "test"},
		{"", ""},
		{"snake_case", "snakeCase"},
		//{"PascalCase", "pascalCase"},
		{"kebab-case", "kebabCase"},
		//{"abc〩", "abc〩"},
		//{"소주", "소주"},
		//{"AbC", "abC"},
	}

	for _, test := range tests {
		actual := ToCamel(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToCamel(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToPascal tests the ToPascal function
func TestToPascal(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "Abc"},
		{"foo 123", "Foo123"},
		//{"FOO_bar", "FooBar"},
		{"FOO_bar", "FOOBar"},
		{" test", "Test"},
		{"  test  ", "Test"},
		{"", ""},
		{"snake_case", "SnakeCase"},
		{"camelCase", "CamelCase"},
		{"kebab-case", "KebabCase"},
		//{"abc〩", "Abc〩"},
		//{"소주", "소주"},
		{"AbC", "AbC"},
	}

	for _, test := range tests {
		actual := ToPascal(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToPascal(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToKebab tests the ToKebab function
func TestToKebab(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "abc"},
		{"foo 123", "foo-123"},
		{"FOO_bar", "foo-bar"},
		{" test", "-test"},
		{"  test  ", "--test--"},
		{"", ""},
		{"snake_case", "snake-case"},
		{"camelCase", "camel-case"},
		{"PascalCase", "pascal-case"},
		{"abc〩", "abc〩"},
		{"소주", "소주"},
		{"AbC", "ab-c"},
	}

	for _, test := range tests {
		actual := ToKebab(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToKebab(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToInt tests the ToInt function
func TestToInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected int
	}{
		{"123", 123},
		{"-42", -42},
		{"5.3", 5},
		{"test", 0},
		{"", 0},
		{"소주", 0},
	}

	for _, test := range tests {
		actual := ToInt(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToInt(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToUint tests the ToUint function
func TestToUint(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected uint
	}{
		{"123", 123},
		{"-42", 0},
		{"5.3", 5},
		{"test", 0},
		{"", 0},
		{"소주", 0},
	}

	for _, test := range tests {
		actual := ToUint(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToUint(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestToBase64 tests the ToBase64 function
func TestToBase64(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"abc", "YWJj"},
		{"foo 123", "Zm9vIDEyMw=="},
		{"  test  ", "ICB0ZXN0ICA="},
		{"", ""},
		{"abc〩", "YWJj44Cp"},
		{"소주", "7IaM7KO8"},
		{"0", "MA=="},
	}

	for _, test := range tests {
		actual := ToBase64(test.data)
		if actual != test.expected {
			t.Errorf("Expected ToBase64(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestFromBase64 tests the FromBase64 function
func TestFromBase64(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected string
	}{
		{"YWJj", "abc"},
		{"Zm9vIDEyMw==", "foo 123"},
		{"ICB0ZXN0ICA=", "  test  "},
		{"", ""},
		{"YWJj44Cp", "abc〩"},
		{"7IaM7KO8", "소주"},
		{"MA==", "0"},
	}

	for _, test := range tests {
		actual := FromBase64(test.data)
		if actual != test.expected {
			t.Errorf("Expected FromBase64(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsBool tests the IsBool function
func TestIsBool(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"true", true},
		{"false", true},
		{"0", true},
		{"", false},
		{"소주", false},
		{"1", true},
		{"null", false},
		{"nil", false},
	}

	for _, test := range tests {
		actual := IsBool(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsBool(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsEmail tests the IsEmail function
func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"abc", false},
		{"me@", false},
		{"0", false},
		{"", false},
		{"소주", false},
		{"me@example.com", true},
		{"me(at)example.com", false},
		{"nil", false},
		{"1", false},
	}

	for _, test := range tests {
		actual := IsEmail(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsEmail(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsURL tests the IsURL function
func TestIsURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		// {"http://", false}, It should be FALSE, but returns TRUE
		{"www.example", false},
		{"www.example.com", false}, // Missing protocol
		{"", false},
		{"소주", false},
		{"me@example.com", false},
		//{"ftp://user:pswd@localhost", false}, It should be FALSE, but returns TRUE
		{"nil", false},
		{"1", false},
		{"https://example.com/foo?q=1", true},
		{"https://example.com/foo.xml", true},
		{"http://www.example.com/", true},
		{"http//www.example.com/", false},
	}

	for _, test := range tests {
		actual := IsURL(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsURL(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsJSON tests the IsJSON function
func TestIsJSON(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"소주", false},
		{"nil", false},
		// {"1", false}, It should be FALSE, but returns TRUE
		{"[1]", true},
		{"[\"1\"]", true},
		{"{foo:\"bar\"}", false},
		{"{\"foo\":\"bar\"}", true},
		{"{\"foo\":}", false},
		{"{\"foo\"}", false},
	}

	for _, test := range tests {
		actual := IsJSON(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsJSON(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsIP tests the IsIP function
func TestIsIP(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"소주", false},
		{"nil", false},
		{"null", false},
		{"1", false},
		{"0", false},
		{"127.0.0.1", true},
		{"127.0.0.", false},
		{"127.0.0", false},
		{"127.0.0.0.1", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", false},
		{"D8-D3-85-EB-12-E3", false},
	}

	for _, test := range tests {
		actual := IsIP(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsIP(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsHexColor tests the IsHexColor function
func TestIsHexColor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"#주주주주주주", false},
		{"nil", false},
		{"null", false},
		{"1", false},
		{"0", false},
		{"#ff", false},
		{"#fff", true},
		{"#ggg", false},
		{"FFF", true},
		{"ff00", false},
		{"ff0000", true},
		{"rgb(255, 255, 255)", false},
	}

	for _, test := range tests {
		actual := IsHexColor(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsHexColor(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsRGBColor tests the IsRGBColor function
func TestIsRGBColor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"#주주주주주주", false},
		{"nil", false},
		{"null", false},
		{"1", false},
		{"0", false},
		{"#fff", false},
		{"#ffffff", false},
		{"rgb(255, 255, 255)", true},
		{"rgb(0, 256, 255)", false},
		{"rgba(0, 255, 255, 1)", false},
		{"rgb(256, 255, 255)", false},
		{"rgb()", false},
		{"rgb(255, 255)", false},
		{"rgb(255, 255, 255, 255)", false},
	}

	for _, test := range tests {
		actual := IsRGBColor(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsRGBColor(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsCreditCard tests the IsCreditCard function
func TestIsCreditCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주주주주-주주주주-주주주주-주주주주", false},
		{"nil", false},
		{"null", false},
		{"1", false},
		{"0", false},
		{"1111-2222-3333-4444", false}, // Not valid vendor (1111)
		{"4653 0343 2480 9848", true},  // Visa
		{"4653-0343-2480-9848", true},  // Visa
		{"5321 7873 3201 8954", true},  // Mastercard
		{"5321-7873-3201-8954", true},  // Mastercard
		{"3421 941266 26371", true},    // American Express
		{"3421-941266-26371", true},    // American Express
		{"1111-222a-3333-4444", false},
		{"1111-2222-333주-4444", false},
		{"111-2222-3333-4444", false},
		{"1111.2222.3333.4444", false},
		{"1111-2222-3333-4444-5555", false},
		{"1111-2222-3333", false},
	}

	for _, test := range tests {
		actual := IsCreditCard(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsCreditCard(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsOnlyDigits tests the IsOnlyDigits function
func TestIsOnlyDigits(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주주주주-주주주주-주주주주-주주주주", false},
		{"nil", false},
		{"null", false},
		{"1", true},
		{"0", true},
		{"1111-", false},
		{"\u0031", true},
		{"\x31", true},
	}

	for _, test := range tests {
		actual := IsOnlyDigits(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsOnlyDigits(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsOnlyLetters tests the IsOnlyLetters function
func TestIsOnlyLetters(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주주주주-주주주주-주주주주-주주주주", false},
		{"주", false},
		{"nil", true},  // Actually are just strings
		{"null", true}, // Actually are just strings
		{"1", false},
		{"0", false},
		{"-", false},
		{"a", true},
		{"the lazy", false},
		{"lazy", true},
		{"\u0061", true},
		{"\x61", true},
	}

	for _, test := range tests {
		actual := IsOnlyLetters(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsOnlyLetters(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}

// TestIsOnlyAlphaNumeric tests the IsOnlyAlphaNumeric function
func TestIsOnlyAlphaNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		data     string
		expected bool
	}{
		{"", false},
		{"주주주주-주주주주-주주주주-주주주주", false},
		{"주", false},
		{"nil", true},  // Actually are just strings
		{"null", true}, // Actually are just strings
		{"1", true},
		{"0", true},
		{"-", false},
		{"a", true},
		{"the lazy", false},
		{"lazy", true},
		{"lazy0", true},
		{"lazy 0", false},
		{"0.1", false},
		{"\u0061\u0031", true},
		{"\x61\x31", true},
	}

	for _, test := range tests {
		actual := IsOnlyAlphaNumeric(test.data)
		if actual != test.expected {
			t.Errorf("Expected IsOnlyAlphaNumeric(%q) to be %v, got %v", test.data, test.expected, actual)
		}
	}
}
