// Package bits provides bit-level operations on byte slices.
package bits

// Test reports whether the bit at position pos in buf is set.
func Test(buf []byte, pos int) bool {
	return buf[pos/8]&(1<<(pos%8)) != 0
}

// Bytes returns a byte slice large enough to hold the given number
// of bits.
func Bytes(bits int) []byte {
	return make([]byte, (bits+7)/8)
}
