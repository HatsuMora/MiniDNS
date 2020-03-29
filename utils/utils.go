package utils

import (
	"encoding/binary"
	"fmt"
)

func FormatAddress(ip string, port int) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

func IntToByteArray(number int, size int) []byte {
	intArray := make([]byte, size)
	binary.BigEndian.PutUint16(intArray, uint16(number))
	return intArray
}
