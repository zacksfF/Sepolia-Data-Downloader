package bchain

import "encoding/binary"

// Int64Bytes converts the integer `i` an `int64` into a byte array
// using Big Endian encoding
func Int64Bytes(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}
