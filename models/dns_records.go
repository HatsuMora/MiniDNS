package models

import (
	"hatsumora.com/MiniDNS/utils"
	"strconv"
	"strings"
)

type NSRecord struct {
	Host string `json:"host"`
}

type ARecord struct {
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Value string `json:"value"`
}

func (record *ARecord) ToByteArray() []byte {
	byteArray := []byte{0xc0, 0x0c, 0, 1, 0, 1}
	byteArray = append(byteArray, utils.IntToByteArray(record.TTL, 4)...)
	byteArray = append(byteArray, byte(0), byte(4))

	for _, part := range strings.Split(record.Value, ".") {
		n, _ := strconv.Atoi(part)
		byteArray = append(byteArray, byte(n))
	}

	return byteArray
}

func (record *NSRecord) ToByteArray() []byte {
	return []byte{0}
}
