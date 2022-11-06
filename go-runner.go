package main

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
)

func main() {

	rand.Seed(time.Now().UnixNano())

	files, err := ioutil.ReadDir("cmpgrounds-pics-jpg")
	if err != nil {
		log.Fatal(err)
	}
	length := len(files)
	randIndex := rand.Intn(length)

	// for idx, file := range files {
	// 	fmt.Println(idx)
	// 	fmt.Println(file.Name())
	// }

	fmt.Println(randIndex)
	fmt.Printf("cmpgrounds-pics-jpgs/%s\n", files[randIndex].Name())

	image1, err := os.Open("cmpgrounds-pics-jpg/1.jpg") // *os.File
	if err != nil {
		log.Println("image1")
		log.Fatal(err)
	}

	first, err := jpeg.Decode(image1) // image.Image
	if err != nil {
		log.Println("first")
		log.Fatal(err)
	}

	defer image1.Close()

	image2, err := os.Open("response.png")
	if err != nil {
		log.Println("image2")
		log.Fatal(err)
	}

	second, err := png.Decode(image2)
	if err != nil {
		log.Println("second")
		log.Fatal(err)
	}

	defer image2.Close()

	offset := image.Pt(750, 750)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.Point{}, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.Point{}, draw.Over)

	third, err := os.Create("results.jpg")
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
