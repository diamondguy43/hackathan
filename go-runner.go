package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	image1, err := os.Open("cmpgrounds-pics-jpg/1.jpg")
	if err != nil {
		log.Println("image1")
		log.Fatal(err)
	}

	first, err := jpeg.Decode(image1)
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

	second, err := jpeg.Decode(image1)
	if err != nil {
		log.Println("second")
		log.Fatal(err)
	}

	defer image2.Close()

	offset := image.Pt(200, 50)
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
