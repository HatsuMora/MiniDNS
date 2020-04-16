package main

import (
	"bytes"
	"fmt"
	"hatsumora.com/MiniDNS/models"
	"hatsumora.com/MiniDNS/utils"
	"log"
	"net"
	"strings"
)

const OpCodeFlag byte = 0b01111000
const maxBufferSize = 512

func handleTraffic(closeChan chan bool) error {
	connection, err := net.ListenPacket("udp", utils.FormatAddress(ip, port))
	if err != nil {
		log.Fatal("Can't listen for packages")
	}
	c := make(chan models.Request)
	defer close(c)

	for {
		go receivePackage(connection, c)
		select {
		case close := <-closeChan:
			if close {
				return connection.Close()
			}

		case req := <-c:
			res := buildResponse(req.Data)
			connection.WriteTo(res, req.Address)
		}
	}
	return connection.Close()
}

func receivePackage(connection net.PacketConn, c chan models.Request) {
	buffer := make([]byte, maxBufferSize)
	n, addr, err := connection.ReadFrom(buffer)
	if err != nil {
		// TODO: Handle error
	}

	fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr.String())
	c <- models.Request{Data: buffer, Address: addr}
}

func buildResponse(data []byte) []byte {
	request := NewDnsRequest(data)
	if request.IsFoulted() {
		return request.ToByteArray()
	}
	var response []byte

	domainQuery, questionTypeIndex := getQuestionDomain(data[12:])
	queryType := getQueryType(data[12+questionTypeIndex+1 : 12+questionTypeIndex+3])
	dnsRecords, nResponses := getDnsRecords(domainQuery, queryType)

	response = append(response, createHeaders(data, nResponses)...)

	response = append(response, buildQuestion(domainQuery, queryType)...)

	response = append(response, dnsRecords...)

	return response
}

func buildQuestion(domainQuery []string, queryType models.QueryType) []byte {
	var question []byte
	for _, domainPart := range domainQuery {
		question = append(question, byte(len(domainPart)))
		question = append(question, []byte(domainPart)...)
	}
	if queryType == models.A {
		question = append(question, 0, 1)
	}

	// Class Internet
	question = append(question, 0, 1)

	return question
}

func createHeaders(data []byte, nRecords int) []byte {
	var headers []byte

	transactionId := data[:2]

	headers = append(headers, transactionId...)
	headers = append(headers, createFlags(data[2:4])...)

	QDCOUNT := data[4:6]
	headers = append(headers, QDCOUNT...)

	// Answer count
	ANCount := utils.IntToByteArray(nRecords, 2)
	headers = append(headers, ANCount...)

	// NSRecord Count
	NSCount := []byte{0, 0}
	headers = append(headers, NSCount...)

	// Additional count
	ARCount := []byte{0, 0}
	headers = append(headers, ARCount...)

	return headers
}

func getDnsRecords(domain []string, queryType models.QueryType) ([]byte, int) {
	zone := getZone(strings.Join(domain, "."))
	if queryType == models.A {
		return zone.ARecordsAsByteArray(), len(zone.ARecords)
	}

	if queryType == models.NS {
		panic("")
	}
	panic("not supported domain query")
}

func getZone(domain string) models.Zone {
	return models.ReadZoneFromFile()
}

func getQueryType(queryTypeBytes []byte) models.QueryType {
	if bytes.Equal(queryTypeBytes, []byte{0, 1}) {
		return models.A
	}
	if bytes.Equal(queryTypeBytes, []byte{0, 2}) {
		return models.NS
	}
	if bytes.Equal(queryTypeBytes, []byte{0, 5}) {
		return models.CNAME
	}
	if bytes.Equal(queryTypeBytes, []byte{0, 12}) {
		return models.MX
	}
	panic("not supported query")
}

func getQuestionDomain(query []byte) ([]string, int) {
	state := 0
	expectedLength := 0
	domainString := ""
	var domainParts []string
	numberOfParts := 0
	n := 0
	for _, payload := range query {
		if payload == 0 {
			domainParts = append(domainParts, domainString)
			break
		}
		if state == 1 {
			domainString += string(payload)
			numberOfParts++
			if numberOfParts == expectedLength {
				domainParts = append(domainParts, domainString)
				domainString = ""
				state = 0
				numberOfParts = 0
			}
		} else {
			state = 1
			expectedLength = int(payload)
		}
		n++
	}
	return domainParts, n
}

func createFlags(flags []byte) []byte {

	QR := byte(0b10000000)
	OpCode := flags[0] & OpCodeFlag
	AA := byte(0b00000100)
	TC := byte(0b00000000)
	RD := byte(0b00000000)
	RA := byte(0b00000000)
	Z := byte(0b00000000)
	RCODE := byte(0b00000000)

	return []byte{QR | AA | OpCode | TC | RD, RA | Z | RCODE}
}
