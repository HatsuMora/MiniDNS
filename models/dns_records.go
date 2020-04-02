package models

import (
	"hatsumora.com/MiniDNS/utils"
	"strconv"
	"strings"
)

/* For future reference:

They must start with a letter, end with a letter or digit, and have as interior
characters only letters, digits, and hyphen.

For all parts of the DNS that are part of the official protocol, all
comparisons between character strings (e.g., labels, domain names, etc.)
are done in a case-insensitive manner

When data enters the domain system, its original case should be
preserved whenever possible

2.3.4. Size limits
labels          63 octets or less
names           255 octets or less
TTL             positive values of a signed 32 bit number.
UDP messages    512 octets or less
*/

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
