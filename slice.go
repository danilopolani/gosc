// Package gosc is an helper package for Go, written to be user friendly with alias and inspired by Lodash.
package gosc

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

// MapString a new slice containing the results of applying the function f to each string in the original slice.
func MapString(s []string, f func(string) string) []string {
	sm := make([]string, len(s))
	for i, v := range s {
		sm[i] = f(v)
	}
	return sm
}

// MapInt a new slice containing the results of applying the function f to each int in the original slice.
func MapInt(s []int, f func(int) int) []int {
	sm := make([]int, len(s))
	for i, v := range s {
		sm[i] = f(v)
	}
	return sm
}

// MapFloat a new slice containing the results of applying the function f to each float64 in the original slice.
func MapFloat(s []float64, f func(float64) float64) []float64 {
	sm := make([]float64, len(s))
	for i, v := range s {
		sm[i] = f(v)
	}
	return sm
}

// FilterString returns a new slice containing all strings in the slice that satisfy the predicate f.
func FilterString(s []string, f func(string) bool) []string {
	sf := make([]string, 0)
	for _, v := range s {
		if f(v) {
			sf = append(sf, v)
		}
	}

	return sf
}

// FilterInt returns a new slice containing all ints in the slice that satisfy the predicate f.
func FilterInt(s []int, f func(int) bool) []int {
	sf := make([]int, 0)
	for _, v := range s {
		if f(v) {
			sf = append(sf, v)
		}
	}

	return sf
}

// FilterFloat returns a new slice containing all float64s in the slice that satisfy the predicate f.
func FilterFloat(s []float64, f func(float64) bool) []float64 {
	sf := make([]float64, 0)
	for _, v := range s {
		if f(v) {
			sf = append(sf, v)
		}
	}

	return sf
}

// AllString returns true if all of the strings in the slice satisfy the predicate f
func AllString(s []string, f func(string) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// AllInt returns true if all of the ints in the slice satisfy the predicate f
func AllInt(s []int, f func(int) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// AllFloat returns true if all of the float64s in the slice satisfy the predicate f
func AllFloat(s []float64, f func(float64) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// EveryString returns true if all of the strings in the slice satisfy the predicate f
func EveryString(s []string, f func(string) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// EveryInt returns true if all of the ints in the slice satisfy the predicate f
func EveryInt(s []int, f func(int) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// EveryFloat returns true if all of the float64s in the slice satisfy the predicate f
func EveryFloat(s []float64, f func(float64) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// AnyString returns true if one of the strings in the slice satisfies the predicate f
func AnyString(s []string, f func(string) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyInt returns true if one of the ints in the slice satisfies the predicate f
func AnyInt(s []int, f func(int) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyFloat returns true if one of the float64s in the slice satisfies the predicate f
func AnyFloat(s []float64, f func(float64) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeString returns true if one of the strings in the slice satisfies the predicate f
func SomeString(s []string, f func(string) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeInt returns true if one of the ints in the slice satisfies the predicate f
func SomeInt(s []int, f func(int) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeFloat returns true if one of the float64s in the slice satisfies the predicate f
func SomeFloat(s []float64, f func(float64) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// Index returns the index of an element in a slice or -1 if not found
func Index(s interface{}, t interface{}) int {
	// Retrieve slices
	sl := reflect.ValueOf(s).Elem()
	if sl.Len() == 0 {
		return -1
	}

	for i := 0; i < sl.Len(); i++ {
		if sl.Index(i).Interface() == t {
			return i
		}
	}

	return -1
}

// Indexi returns the index of a string in a slice or -1 if not found. Case Insentive.
func Indexi(s []string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return -1
	}

	for i, v := range s {
		if strings.ToLower(v) == strings.ToLower(t) {
			return i
		}
	}

	return -1
}

// Delete an item from a slice
func Delete(s interface{}, i int) {
	// Retrieve slice
	sl := reflect.ValueOf(s).Elem()

	// Return if index not present or empty slice
	if i > sl.Len()-1 || sl.Len() == 0 || i < 0 {
		return
	}

	// Exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		*s.(*[]string) = append(sli[:i], sli[i+1:]...)
	case *[]int:
		sli := sl.Interface().([]int)
		*s.(*[]int) = append(sli[:i], sli[i+1:]...)
	case *[]float64:
		sli := sl.Interface().([]float64)
		*s.(*[]float64) = append(sli[:i], sli[i+1:]...)
	default:
		return
	}
}

// Rsort reverses the order (desc) of a slice
func Rsort(s interface{}) {
	// Retrieve slice
	sl := reflect.ValueOf(s).Elem()

	// Handle types or exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		sort.Sort(sort.Reverse(sort.StringSlice(sli)))
		*s.(*[]string) = sli
	case *[]int:
		sli := sl.Interface().([]int)
		sort.Sort(sort.Reverse(sort.IntSlice(sli)))
		*s.(*[]int) = sli
	case *[]float64:
		sli := sl.Interface().([]float64)
		sort.Sort(sort.Reverse(sort.Float64Slice(sli)))
		*s.(*[]float64) = sli
	default:
		return
	}
}

// ReverseSort is an alias of Rsort
func ReverseSort(s interface{}) {
	// Retrieve slice
	sl := reflect.ValueOf(s).Elem()

	// Handle types or exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		sort.Sort(sort.Reverse(sort.StringSlice(sli)))
		*s.(*[]string) = sli
	case *[]int:
		sli := sl.Interface().([]int)
		sort.Sort(sort.Reverse(sort.IntSlice(sli)))
		*s.(*[]int) = sli
	case *[]float64:
		sli := sl.Interface().([]float64)
		sort.Sort(sort.Reverse(sort.Float64Slice(sli)))
		*s.(*[]float64) = sli
	default:
		return
	}
}

// EqSlices returns true if two slices are equal (not in-depth, for that use reflect.DeepEqual)
func EqSlices(a, b interface{}) bool {
	// Retrieve slices
	sl1 := reflect.ValueOf(a).Elem()
	sl2 := reflect.ValueOf(b).Elem()

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if sl1.Len() != sl2.Len() {
		return false
	}

	if fmt.Sprintf("%T", sl1.Interface()) != fmt.Sprintf("%T", sl1.Interface()) {
		return false
	}

	for i := 0; i < sl1.Len(); i++ {
		if sl1.Index(i).Interface() != sl2.Index(i).Interface() {
			return false
		}
	}

	return true
}

// SliceRand assign a random value from a string to the "r" arg
func SliceRand(s interface{}, r interface{}) {
	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// Retrieve slice
	sl := reflect.ValueOf(s).Elem()

	// Handle types or exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		*r.(*string) = sli[rand.Intn(len(sli))]
	case *[]int:
		sli := sl.Interface().([]int)
		*r.(*int) = sli[rand.Intn(len(sli))]
	case *[]float64:
		sli := sl.Interface().([]float64)
		*r.(*float64) = sli[rand.Intn(len(sli))]
	default:
		return
	}
}

// SliceRandString returns a random value from a string slice
func SliceRandString(s []string) string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return s[rand.Intn(len(s))]
}

// SliceRandS is an alias of SliceRandString
func SliceRandS(s []string) string {
	return SliceRandString(s)
}

// SliceRandInt returns a random value from an int slice
func SliceRandInt(s []int) int {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return s[rand.Intn(len(s))]
}

// SliceRandI is an alias of SliceRandInt
func SliceRandI(s []int) int {
	return SliceRandInt(s)
}

// SliceRandFloat returns a random value from a float64 slice
func SliceRandFloat(s []float64) float64 {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return s[rand.Intn(len(s))]
}

// SliceRandF is an alias of SliceRandFloat
func SliceRandF(s []float64) float64 {
	return SliceRandFloat(s)
}

// InSlice returns a boolean if the value is in the slice
// v = value to find
// s = slice
func InSlice(v interface{}, s interface{}) bool {
	// Check if is passed as pointer
	valueOf := reflect.ValueOf(s)
	if valueOf.Kind() != reflect.Ptr {
		panic("Non-pointer slice provided.")
	}

	// Retrieve slice
	sl := valueOf.Elem()

	// Invoke correct function
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return inIntSlice(v.(int), sl)
	case reflect.Float64:
		return inFloat64Slice(v.(float64), sl)
	case reflect.String:
		return inStringSlice(v.(string), sl)
	case reflect.Bool:
		return inBoolSlice(v.(bool), sl)
	default:
		return false
	}
}

func inIntSlice(v int, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		if v == s.Index(i).Interface().(int) {
			return true
		}
	}

	return false
}

func inFloat64Slice(v float64, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		if v == s.Index(i).Interface().(float64) {
			return true
		}
	}

	return false
}

func inStringSlice(v string, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		if v == s.Index(i).Interface().(string) {
			return true
		}
	}

	return false
}

func inBoolSlice(v bool, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		if v == s.Index(i).Interface().(bool) {
			return true
		}
	}

	return false
}
