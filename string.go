package gosc

import (
	"bytes"
	cryrand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

// ToBytes converts a string to bytes slice
func ToBytes(s string) []byte {
	return []byte(s)
}

// ByteToString converts a bytes slice to a string
func ByteToString(b []byte) string {
	return string(b[:])
}

// Rstring returns the string reversed
func Rstring(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ReverseString is an alias of Rstring
func ReverseString(s string) string {
	return Rstring(s)
}

// LcFirst returns a string with the first character lowercased
func LcFirst(s string) string {
	if s == "" {
		return ""
	}

	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// LowerFirst is an alias of LcFirst
func LowerFirst(s string) string {
	return LcFirst(s)
}

// UpperFirst returns a string with the first character uppercased
func UcFirst(s string) string {
	if s == "" {
		return ""
	}

	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

// UpperFirst is an alias of UcFirst
func UpperFirst(s string) string {
	return UcFirst(s)
}

// ------------------
// "To" section
// ------------------

// ToSnake converts a string to snake_case
func ToSnake(s string) string {
	s = strings.Replace(strings.Replace(s, "-", "_", -1), " ", "_", -1)
	runes := []rune(s)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// ToSnakeCase is an alias of ToSnake
func ToSnakeCase(s string) string {
	return ToSnake(s)
}

// ToCamel converts a string to camelCase
func ToCamel(s string) string {
	byteSrc := []byte(s)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		val = bytes.ToLower(val)

		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}

	return string(bytes.Join(chunks, nil))
}

// ToCamelCase is an alias of ToCamel
func ToCamelCase(s string) string {
	return ToCamel(s)
}

// ToPascal converts a string to PascalCase
func ToPascal(s string) string {
	byteSrc := []byte(s)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}

	return string(bytes.Join(chunks, nil))
}

// ToPascalCase is an alias of ToPascal
func ToPascalCase(s string) string {
	return ToPascal(s)
}

// ToKebab converts a string to kebab-case
func ToKebab(s string) string {
	s = strings.Replace(strings.Replace(s, "_", "-", -1), " ", "-", -1)
	runes := []rune(s)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '-')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// ToKebabCase is an alias of ToKebab
func ToKebabCase(s string) string {
	return ToKebab(s)
}

// ToInt returns an int from a string
func ToInt(s string) int {
	if strings.Contains(s, ".") {
		s = strings.Split(s, ".")[0]
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return v
}

// ToUint returns a uint from a string
func ToUint(s string) uint {
	if strings.Contains(s, ".") {
		s = strings.Split(s, ".")[0]
	}

	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return uint(0)
	}

	return uint(v)
}

// ToBase64 encodes a string in base64
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// FromBase64 decodes a string from base64
func FromBase64(s string) string {
	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}

	return string(str)
}

// ------------------
// "Is" section
// ------------------

// IsBool checks if a string is a boolean
func IsBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsEmail returns true if the provided string is a valid email address
func IsEmail(s string) bool {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegexp.MatchString(s)
}

// IsURL returns true if the provided string is a valid url
func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// IsJSON returns true if the provided string is a valid JSON document
func IsJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsIP returns true if the provided string is a valid IPv4
func IsIP(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && strings.Count(s, ".") == 3
}

// IsHexColor returns true if the provided string is a valid HEX color
func IsHexColor(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 3 && len(s) != 6 {
		return false
	}

	for _, c := range s {
		if ('F' < c || c < 'A') && ('f' < c || c < 'a') && ('9' < c || c < '0') {
			return false
		}
	}

	return true
}

// IsRGBColor returns true if the provided string is a valid RGB color
func IsRGBColor(s string) bool {
	if s == "" || len(s) < 10 {
		return false
	}

	if s[0:4] != "rgb(" || s[len(s)-1] != ')' {
		return false
	}

	s = s[4 : len(s)-1]
	s = strings.TrimSpace(s)

	if strings.Count(s, ",") != 2 {
		return false
	}

	for _, p := range strings.Split(s, ",") {
		if len(p) > 1 && p[0] == '0' {
			return false
		}

		p = strings.TrimSpace(p)
		if i, e := strconv.Atoi(p); (255 < i || i < 0) || e != nil {
			return false
		}
	}

	return true
}

// IsRGB is an alias of IsRGBColor
func IsRGB(s string) bool {
	return IsRGBColor(s)
}

// IsCreditCard returns true if the provided string is a valid credit card
func IsCreditCard(s string) bool {
	rxCreditCard := regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$")
	r, _ := regexp.Compile("[^0-9]+")
	sanitized := r.ReplaceAll([]byte(s), []byte(""))

	if !rxCreditCard.MatchString(string(sanitized)) {
		return false
	}

	var sum int
	var digit string
	var tmpNum int
	var shouldDouble bool

	for i := len(sanitized) - 1; i >= 0; i-- {
		digit = string(sanitized[i:(i + 1)])
		tmpNum = ToInt(digit)
		if shouldDouble {
			tmpNum *= 2
			if tmpNum >= 10 {
				sum += ((tmpNum % 10) + 1)
			} else {
				sum += tmpNum
			}
		} else {
			sum += tmpNum
		}
		shouldDouble = !shouldDouble
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

// IsOnlyDigits returns true if the provided string is composed only by numbers
func IsOnlyDigits(s string) bool {
	r := regexp.MustCompile("^[0-9]$")
	return r.MatchString(s)
}

// IsOnlyNumbers is an alias of IsOnlyDigits
func IsOnlyNumbers(s string) bool {
	return IsOnlyDigits(s)
}

// IsOnlyLetters returns true if the provided string is composed only by letters
func IsOnlyLetters(s string) bool {
	r := regexp.MustCompile("^[A-Za-z]+$")
	return r.MatchString(s)
}

// IsOnlyAlpha is an alias of IsOnlyLetters
func IsOnlyAlpha(s string) bool {
	return IsOnlyLetters(s)
}

// IsOnlyAlphaNumeric returns true if the provided string is composed only by letters and numbers
func IsOnlyAlphaNumeric(s string) bool {
	r := regexp.MustCompile("^[A-Za-z0-9]+$")
	return r.MatchString(s)
}

// IsOnlyAlphaDigits is an alias of IsOnlyAlphaNumeric
func IsOnlyAlphaDigits(s string) bool {
	return IsOnlyAlphaNumeric(s)
}

// IsOnlyLettersNumbers is an alias of IsOnlyAlphaNumeric
func IsOnlyLettersNumbers(s string) bool {
	return IsOnlyAlphaNumeric(s)
}

// IsOnlyAlphaNum is an alias of IsOnlyAlphaNumeric
func IsOnlyAlphaNum(s string) bool {
	return IsOnlyAlphaNumeric(s)
}

// ------------------
// "Generator" section
// ------------------

// Unq returns a unique token based on current time and crypted in sha256
func Uniq() string {
	t := strconv.Itoa(int(time.Now().UnixNano()))
	return fmt.Sprintf("%x", sha256.Sum256([]byte("gosc_"+t)))
}

// Unique is an alias of Uniq
func Unique() string {
	return Uniq()
}

// StrRand returns a random string
func StrRand(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandStr is an alias of StrRand
func RandStr(n int) string {
	return StrRand(n)
}

// RandomString is an alias of StrRand
func RandomString(n int) string {
	return StrRand(n)
}

// UUID generates a random UUID according to RFC 4122
func UUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(cryrand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
