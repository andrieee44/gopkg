// Package bits provides bit-level operations.
package bitops

import "golang.org/x/exp/constraints"

// Test reports whether the bit at position pos in buf is set.
func Test[T constraints.Integer](buf []byte, pos T) bool {
	return buf[pos/8]&(1<<(pos%8)) != 0
}

// Bytes returns a byte slice large enough to hold the given number
// of bits.
func Bytes[T constraints.Integer](bits T) []byte {
	return make([]byte, (bits+7)/8)
}

// OverflowsSigned reports whether value cannot be represented
// in a signed integer type using the given number of bits.
func OverflowsSigned[T constraints.Signed](bits, value T) bool {
	return value < -(1<<(bits-1)) || value > (1<<(bits-1))-1
}

// OverflowsUnsigned reports whether value cannot be represented
// in an unsigned integer type using the given number of bits.
func OverflowsUnsigned[T constraints.Unsigned](bits, value T) bool {
	return value > (1<<bits)-1
}
