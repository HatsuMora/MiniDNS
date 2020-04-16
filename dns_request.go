package main

//                                    1  1  1  1  1  1
//      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                      ID                       |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    QDCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ANCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    NSCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ARCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+

type DnsRequest struct {
	Id      [2]byte
	QR      bool // 0 = false = request, 1 = true = response
	Opcode  uint8
	AA      bool
	TC      bool
	RD      bool
	RA      bool
	Z       uint8
	RCODE   uint8
	QDCount uint16
	ANCount uint16
	ARCount uint16
}

func NewDnsRequest(data []byte) *DnsRequest {
	if data == nil {
		panic("Error: data cannot be nil")
	}
	if len(data) < 12 {
		panic("Error: Length of data must be at least 12 bytes long")
	}
	if len(data) > maxBufferSize {
		panic("This server does not support datagrams longer than " + string(maxBufferSize) + " bytes")
	}

	req := &DnsRequest{}
	req.Id = [2]byte{data[0], data[1]}
	req.QR = false
	req.Opcode = data[2] & OpCodeFlag >> 3
	req.AA = false
	req.TC = false
	req.RD = (byte(0b00000001) & data[2]) == 1
	req.RA = false
	req.RCODE = 0
	return req
}

func (req DnsRequest) ToByteArray() []byte {
	return nil
}

func (req DnsRequest) IsFoulted() bool {
	if req.RCODE != 0 {
		return true
	}
	return false
}
