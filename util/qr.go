package util

import (
	"log"
	"math/rand"

	b64 "encoding/base64"
	"encoding/json"

	qrcode "github.com/skip2/go-qrcode"
)

// QRKey: Key Information Embedded into QR Code
type QRKey struct {
	// UUID       string `json:"uuid"`
	StartDate  string `json:"startdate"`
	EndDate    string `json:"enddate"`
	Location   string `json:"location"`
	RandomBits int    `json:"randombits"`
}

func QRKeyGen(sDate, eDate, location string) []byte {

	key := &QRKey{
		StartDate:  sDate,
		EndDate:    eDate,
		Location:   location,
		RandomBits: rand.Int(),
	}

	// Debug to print key
	// fmt.Println(key)

	jsonData, err := json.Marshal(key)
	if err != nil {
		log.Fatal(err)
	}

	obfuscatedData := b64.StdEncoding.EncodeToString(jsonData)

	return QRCodeGen(obfuscatedData)
}

// QRCodeGen: function response for generating the actual QR Code
func QRCodeGen(content string) []byte {
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	return png
}