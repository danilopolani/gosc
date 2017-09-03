// Package gosc is an helper package for Go, written to be user friendly with alias and inspired by Lodash.
package gosc

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// IsInt checks if a string is an integer
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// IsFloat checks if a string is a float number
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Utoa transforms a uint into a string
func Utoa(u uint) string {
	return fmt.Sprint(u)
}

// Rand returns a random int from the given range
func Rand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
