package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type QueryType string

const (
	A     QueryType = "a"
	CNAME QueryType = "cname"
	NS    QueryType = "ns"
	MX    QueryType = "mx"
	SRV   QueryType = "srv"
)

//3.2.2. TYPE values
//
//TYPE fields are used in resource records.  Note that these types are a
//subset of QTYPEs.
//
//TYPE            value and meaning
//
//A               1 a host address
//NS              2 an authoritative name server
//MD              3 a mail destination (Obsolete - use MX)
//MF              4 a mail forwarder (Obsolete - use MX)
//CNAME           5 the canonical name for an alias
//SOA             6 marks the start of a zone of authority
//MB              7 a mailbox domain name (EXPERIMENTAL)
//MG              8 a mail group member (EXPERIMENTAL)
//MR              9 a mail rename domain name (EXPERIMENTAL)
//NULL            10 a null RR (EXPERIMENTAL)
//WKS             11 a well known service description
//PTR             12 a domain name pointer
//HINFO           13 host information
//MINFO           14 mailbox or mail list information
//MX              15 mail exchange
//TXT             16 text strings

type SOA struct {
	Mname   string `json:"mname"`
	Rname   string `json:"rname"`
	Serial  string `json:"serial"`
	Refresh int    `json:"refresh"`
	Retry   int    `json:"retry"`
	Expire  int    `json:"expire"`
	Minimum int    `json:"minimum"`
}

type Zone struct {
	Origin    string     `json:"$origin"`
	Ttl       int        `json:"$ttl"`
	SOA       SOA        `json:"soa"`
	NsRecords []NSRecord `json:"ns"`
	ARecords  []ARecord  `json:"a"`
}

type Request struct {
	Data    []byte
	Address net.Addr
}

func (zone *Zone) ARecordsAsByteArray() []byte {
	var records []byte
	for _, record := range zone.ARecords {
		records = append(records, record.ToByteArray()...)
	}
	return records
}

func ReadZoneFromFile() Zone {
	// Open our jsonFile
	jsonFile, err := os.Open("zones/hatsumora.com.zone")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened hatsumora.com.zone")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var zone Zone

	json.Unmarshal(byteValue, &zone)

	return zone
}
