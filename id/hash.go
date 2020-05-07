package id

import (
	"strconv"
)

//HashCode 取hash值
func HashCode(in string) int32 {
	// Initialize output
	var hash int32
	// Empty string has a hashcode of 0
	if len(in) == 0 {
		return hash
	}
	// Convert string into slice of bytes
	b := []byte(in)
	// Build hash
	for i := range b {
		char := b[i]
		hash = ((hash << 5) - hash) + int32(char)
	}
	return hash
}

//HashString 取hash值，换成串
func HashString(in string) string {
	code := int64(HashCode(in))
	if code < 0 {
		code = ^code + 1
	}
	return strconv.FormatInt(code, 36)
}
