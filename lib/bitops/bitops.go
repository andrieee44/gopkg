// Package bitops provides bit-level operations.
package bitops

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Test reports whether the bit at position pos in buf is set.
//
// Panics if pos is negative or if pos/8 is out of bounds for buf.
func Test[T constraints.Integer](buf []byte, pos T) bool {
	if pos < 0 {
		panic(fmt.Sprintf("bitops.Test: pos: %d is negative; must be >= 0", pos))
	}

	if int(pos/8) >= len(buf) {
		panic(fmt.Sprintf("bitops.Test: pos: %d is out of bounds; buf holds %d bits", pos, len(buf)*8))
	}

	return buf[pos/8]&(1<<(pos%8)) != 0
}

// Bytes returns a byte slice large enough to hold the given number
// of bits.
//
// Panics if bits is negative.
func Bytes[T constraints.Integer](bits T) []byte {
	if bits < 0 {
		panic(fmt.Sprintf("bitops.Bytes: bits: %d is invalid; must be >= 0", bits))
	}

	return make([]byte, (bits+7)/8)
}

// OverflowsSigned reports whether value cannot be represented
// using the specified number of bits in two's complement form.
//
// The check assumes signed binary encoding: value overflows
// if it is less than -2^(bits-1) or greater than or equal to 2^(bits-1).
//
// Panics if bits <= 1 (minimum 1 data bit + 1 sign bit).
// Panics if bits > 64; values wider than int64 are not supported.
func OverflowsSigned[T constraints.Signed](bits, value T) bool {
	var maxN, minN int64

	if bits <= 1 {
		panic(fmt.Sprintf("bitops.OverflowsSigned: bits: %d is invalid; must be > 1 to allow for sign bit", bits))
	}

	if bits > 64 {
		panic(fmt.Sprintf("bitops.OverflowsSigned: bits: %d is invalid; must be <= 64; values wider than int64 are not supported", bits))
	}

	maxN = int64((uint64(1) << (bits - 1)) - 1)
	minN = -int64(1 << (bits - 1))

	return !(int64(value) <= maxN && int64(value) >= minN)
}

// OverflowsUnsigned reports whether value cannot be represented
// using the specified number of bits in unsigned binary form.
//
// The check assumes unsigned encoding: value overflows if it is
// greater than or equal to 2^bits.
//
// Panics if bits == 0.
// Panics if bits > 64; values wider than uint64 are not supported.
func OverflowsUnsigned[T constraints.Unsigned](bits, value T) bool {
	if bits == 0 {
		panic(fmt.Sprintf("bitops.OverflowsUnsigned: bits: %d is invalid; must be > 0", bits))
	}

	if bits > 64 {
		panic(fmt.Sprintf("bitops.OverflowsUnsigned: bits: %d is invalid; must be <= 64; values wider than uint64 are not supported", bits))
	}

	return !(uint64(value) <= ^uint64(0)>>(64-bits))
}
