package util

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

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
	Check(err)

	obfuscatedData := b64.StdEncoding.EncodeToString(jsonData)

	// take QR code
	qrData := QRCodeGen(obfuscatedData)
	err = os.WriteFile("tmp-qr.png", qrData, 0644)
	Check(err)

	// random seed
	rand.Seed(time.Now().UnixNano())

	// list files in cmpgrounds-pics-jpg dir
	files, err := ioutil.ReadDir("cmpgrounds-pics-jpg")
	Check(err)
	length := len(files)
	randIndex := rand.Intn(length)

	file := fmt.Sprintf("cmpgrounds-pics-jpg/%s", files[randIndex].Name())

	im1, err := os.Open(file)
	Check(err)

	bgImg, err := jpeg.Decode(im1)
	Check(err)

	defer im1.Close()

	im2, err := os.Open("tmp-qr.png")
	Check(err)

	overlay, err := png.Decode(im2)
	Check(err)

	defer im2.Close()

	// should be determined by Photo and QR Code Dimensions
	offset := image.Pt(750, 750)
	b := bgImg.Bounds()
	im3 := image.NewRGBA(b)
	draw.Draw(im3, b, bgImg, image.Point{}, draw.Src)
	draw.Draw(im3, overlay.Bounds().Add(offset), overlay, image.Point{}, draw.Over)

	qrkey, err := os.Create("QRKEY.jpg")
	Check(err)

	jpeg.Encode(qrkey, im3, &jpeg.Options{jpeg.DefaultQuality})
	defer qrkey.Close()

	qrkeyb, err := os.ReadFile("QRKEY.jpg")
	Check(err)

	return qrkeyb
}

// QRCodeGen: function response for generating the actual QR Code
func QRCodeGen(content string) []byte {
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	return png
}
