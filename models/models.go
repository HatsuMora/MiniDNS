package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
