package bitops_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/andrieee44/gopkg/lib/bitops"
	"golang.org/x/exp/constraints"
)

func bufStr(buf []byte) string {
	var (
		builder strings.Builder
		i       int
	)

	builder.WriteString("{ ")

	for i = range buf {
		builder.WriteString(fmt.Sprintf("%08b", buf[i]))

		if i != len(buf)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(" }")

	return builder.String()
}

func expectPanic(t *testing.T, format string, args ...any) {
	var r any

	t.Helper()

	r = recover()
	if r == nil {
		t.Errorf(format, args...)
	}
}

func TestTest(t *testing.T) {
	type table struct {
		buf []byte
		pos int
		exp bool
	}

	var (
		tables []table
		test   table
	)

	t.Parallel()

	tables = []table{
		// Single byte
		{[]byte{0b00000001}, 0, true},  // bit 0 set
		{[]byte{0b00000001}, 1, false}, // bit 1 unset
		{[]byte{0b10000000}, 7, true},  // bit 7 set
		{[]byte{0b01000000}, 6, true},  // bit 6 set
		{[]byte{0b00000000}, 3, false}, // bit 3 unset

		// Two bytes
		{[]byte{0b00000000, 0b00000001}, 8, true},  // second byte, bit 0
		{[]byte{0b00000000, 0b00000010}, 9, true},  // second byte, bit 1
		{[]byte{0b00000000, 0b00000000}, 9, false}, // second byte, bit 1 unset
		{[]byte{0b11111111, 0b11111111}, 15, true}, // second byte, bit 7 set

		// Three bytes
		{[]byte{0b00000000, 0b00000000, 0b00010000}, 20, true},  // third byte, bit 4
		{[]byte{0b00000000, 0b00000000, 0b00000000}, 23, false}, // third byte, bit 7 unset

		// Full byte set
		{[]byte{0xFF}, 0, true},
		{[]byte{0xFF}, 7, true},
		{[]byte{0xFF}, 3, true},

		// Edge of buffer
		{[]byte{0x00, 0x00, 0x08}, 19, true}, // third byte, bit 3
	}

	for _, test = range tables {
		if test.exp != bitops.Test(test.buf, test.pos) {
			t.Errorf(
				"got: %t, exp: %t: buf = %s, pos = %d",
				!test.exp,
				test.exp,
				bufStr(test.buf),
				test.pos,
			)
		}
	}
}

func FuzzTest(f *testing.F) {
	f.Add([]byte{}, 0)              // empty buffer, any pos should panic
	f.Add([]byte{0x00}, -1)         // negative pos
	f.Add([]byte{0x00}, 8)          // pos just out of bounds
	f.Add([]byte{0x00}, 1000)       // way out of bounds
	f.Add([]byte{0x00, 0x00}, 15)   // upper edge of valid
	f.Add([]byte{0x00, 0x00}, 16)   // just over valid
	f.Add([]byte{0x00, 0x00}, -100) // deep negative
	f.Add([]byte{0xFF}, 7)          // valid upper bit
	f.Add([]byte{0xFF}, 8)          // just over
	f.Add([]byte{0xFF}, 100)        // far over

	f.Fuzz(func(t *testing.T, buf []byte, pos int) {
		if pos < 0 || pos/8 >= len(buf) {
			defer expectPanic(t, "expected panic: buf = %s, pos = %d", bufStr(buf), pos)
		}

		_ = bitops.Test(buf, pos)
	})
}

func TestBytes(t *testing.T) {
	type table struct {
		bits int
		size int
	}

	var (
		tables []table
		test   table
	)

	t.Parallel()

	tables = []table{
		{0, 0},      // 0 bits = 0 bytes
		{1, 1},      // 1 bit = 1 byte
		{7, 1},      // 7 bits = 1 byte
		{8, 1},      // 8 bits = 1 byte
		{9, 2},      // 9 bits = 2 bytes
		{15, 2},     // 15 bits = 2 bytes
		{16, 2},     // 16 bits = 2 bytes
		{17, 3},     // 17 bits = 3 bytes
		{63, 8},     // 63 bits = 8 bytes
		{64, 8},     // 64 bits = 8 bytes
		{65, 9},     // 65 bits = 9 bytes
		{1000, 125}, // 1000 bits = 125 bytes
	}

	for _, test = range tables {
		if len(bitops.Bytes(test.bits)) != test.size {
			t.Errorf(
				"got: %d, exp: %d: bits = %d",
				len(bitops.Bytes(test.bits)),
				test.size,
				test.bits,
			)
		}
	}
}

func FuzzBytes(f *testing.F) {
	f.Add(-1)            // negative bit count
	f.Add(-100)          // large negative
	f.Add(0)             // zero bits
	f.Add(1)             // minimal non-zero
	f.Add(7)             // just under byte boundary
	f.Add(8)             // exact byte boundary
	f.Add(9)             // just over byte boundary
	f.Add(63)            // one less than 64-bit word
	f.Add(64)            // full 64-bit word
	f.Add(65)            // just over word boundary
	f.Add(255)           // full byte range
	f.Add(256)           // next byte
	f.Add(1023)          // just under 1K bits
	f.Add(1024)          // 1K bits
	f.Add(1025)          // just over 1K
	f.Add(1 << 20)       // upper bound
	f.Add((1 << 20) + 1) // just over upper bound
	f.Add(1 << 30)       // large allocation

	f.Fuzz(func(t *testing.T, bits int) {
		var buf []byte

		if bits < 0 {
			defer expectPanic(t, "expected panic: bits = %d", bits)
		}

		buf = bitops.Bytes(bits)
		if len(buf) != (bits+7)/8 {
			t.Errorf(
				"got: %d, exp: %d: bits = %d",
				len(buf),
				(bits+7)/8,
				bits,
			)
		}
	})
}

func TestOverflowsSigned(t *testing.T) {
	type table[T constraints.Signed] struct {
		bits  T
		value T
		exp   bool
	}

	t.Parallel()

	t.Run("int8", func(t *testing.T) {
		var (
			tests []table[int8]
			test  table[int8]
		)

		tests = []table[int8]{
			// 2-bit signed range [-2, 1]
			{2, -3, true},  // below min
			{2, -2, false}, // min
			{2, 1, false},  // max
			{2, 2, true},   // above max

			// 8-bit signed range [-128, 127]
			{8, -128, false}, // min
			{8, 127, false},  // max
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsSigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("int16", func(t *testing.T) {
		var (
			tests []table[int16]
			test  table[int16]
		)

		tests = []table[int16]{
			// 8-bit signed range [-128, 127]
			{8, -129, true},  // below min
			{8, -128, false}, // min
			{8, 127, false},  // max
			{8, 128, true},   // above max

			// 16-bit signed range [-32768, 32767]
			{16, -32768, false}, // min
			{16, 32767, false},  // max
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsSigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("int32", func(t *testing.T) {
		var (
			tests []table[int32]
			test  table[int32]
		)

		tests = []table[int32]{
			// 32-bit signed range [-2147483648, 2147483647]
			{32, -2147483648, false}, // min
			{32, 2147483647, false},  // max
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsSigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("int64", func(t *testing.T) {
		var (
			tests []table[int64]
			test  table[int64]
		)

		tests = []table[int64]{
			// 64-bit signed range [-9223372036854775808, 9223372036854775807]
			{64, -9223372036854775808, false}, // min
			{64, 9223372036854775807, false},  // max
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsSigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("int", func(t *testing.T) {
		var (
			tests []table[int]
			test  table[int]
		)

		tests = []table[int]{
			// 8-bit signed range [-128, 127]
			{8, -129, true},  // below min
			{8, -128, false}, // min
			{8, 127, false},  // max
			{8, 128, true},   // above max
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsSigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})
}

func FuzzOverflowsSigned(f *testing.F) {
	f.Add(-1, 0)           // negative bits
	f.Add(0, 0)            // zero bits
	f.Add(1, -2)           // underflow for 1-bit signed
	f.Add(1, 0)            // valid for 1-bit signed
	f.Add(1, 1)            // overflow for 1-bit signed
	f.Add(8, -128)         // lower bound for 8-bit signed
	f.Add(8, 127)          // upper bound for 8-bit signed
	f.Add(8, -129)         // underflow for 8-bit signed
	f.Add(8, 128)          // overflow for 8-bit signed
	f.Add(16, -32768)      // lower bound for 16-bit signed
	f.Add(16, 32767)       // upper bound for 16-bit signed
	f.Add(16, -32769)      // underflow for 16-bit signed
	f.Add(16, 32768)       // overflow for 16-bit signed
	f.Add(32, -2147483648) // lower bound for 32-bit signed
	f.Add(32, 2147483647)  // upper bound for 32-bit signed
	f.Add(32, -2147483649) // underflow for 32-bit signed
	f.Add(32, 2147483648)  // overflow for 32-bit signed

	f.Fuzz(func(t *testing.T, bits, value int) {
		if bits <= 1 || bits > 64 {
			defer expectPanic(t, "expected panic: bits = %d, value = %d", bits, value)
		}

		_ = bitops.OverflowsSigned(bits, value)
	})
}

func TestOverflowsUnsigned(t *testing.T) {
	type table[T constraints.Unsigned] struct {
		bits  T
		value T
		exp   bool
	}

	t.Parallel()

	t.Run("uint8", func(t *testing.T) {
		var (
			tests []table[uint8]
			test  table[uint8]
		)

		tests = []table[uint8]{
			// 1-bit unsigned range [0, 1]
			{1, 0, false},
			{1, 1, false},
			{1, 2, true}, // overflow

			// 8-bit unsigned range [0, 255]
			{8, 0, false},
			{8, 255, false},
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsUnsigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("uint16", func(t *testing.T) {
		var (
			tests []table[uint16]
			test  table[uint16]
		)

		tests = []table[uint16]{
			// 8-bit unsigned range [0, 255]
			{8, 256, true},
			{8, 255, false},

			// 16-bit unsigned range [0, 65535]
			{16, 65535, false},
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsUnsigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var (
			tests []table[uint32]
			test  table[uint32]
		)

		tests = []table[uint32]{
			// 32-bit unsigned range [0, 4294967295]
			{32, 4294967295, false},
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsUnsigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("uint64", func(t *testing.T) {
		var (
			tests []table[uint64]
			test  table[uint64]
		)

		tests = []table[uint64]{
			// 64-bit unsigned range [0, 18446744073709551615]
			{64, 0, false},
			{64, 18446744073709551615, false},
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsUnsigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})

	t.Run("uint", func(t *testing.T) {
		var (
			tests []table[uint]
			test  table[uint]
		)

		tests = []table[uint]{
			// 8-bit unsigned range [0, 255]
			{8, 255, false},
			{8, 256, true},
		}

		for _, test = range tests {
			if test.exp != bitops.OverflowsUnsigned(test.bits, test.value) {
				t.Errorf(
					"got: %t, exp: %t: bits = %d, value = %d",
					!test.exp,
					test.exp,
					test.bits,
					test.value,
				)
			}
		}
	})
}

func FuzzOverflowsUnsigned(f *testing.F) {
	f.Add(uint(0), uint(0))                     // zero bits
	f.Add(uint(1), uint(0))                     // fits in 1 bit
	f.Add(uint(1), uint(1))                     // fits in 1 bit
	f.Add(uint(1), uint(2))                     // overflow for 1 bit
	f.Add(uint(2), uint(3))                     // max for 2 bits
	f.Add(uint(2), uint(4))                     // overflow for 2 bits
	f.Add(uint(8), uint(255))                   // max for 8 bits
	f.Add(uint(8), uint(256))                   // overflow for 8 bits
	f.Add(uint(16), uint(65535))                // max for 16 bits
	f.Add(uint(16), uint(65536))                // overflow for 16 bits
	f.Add(uint(32), uint(4294967295))           // max for 32 bits
	f.Add(uint(32), uint(4294967296))           // overflow for 32 bits
	f.Add(uint(64), uint(18446744073709551615)) // max for 64 bits

	f.Fuzz(func(t *testing.T, bits, value uint) {
		if bits == 0 || bits > 64 {
			defer expectPanic(t, "expected panic: bits = %d, value = %d", bits, value)
		}

		_ = bitops.OverflowsUnsigned(bits, value)
	})
}
