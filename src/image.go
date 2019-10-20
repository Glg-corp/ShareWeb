package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	
)

func startCompareImage(path string) (bool, string) {

	// Open file
	file, _ := os.Open(path)
	defer file.Close()

	// get the color mean
	pixelsImage, err := getPixels(path)

	if err != nil {
		panic(err)
	}

	width := len(pixelsImage[0])
	height := len(pixelsImage)

	// calculate the mean
	mean := computeMeanSlow(height, width, 0, 0, pixelsImage)

	hex := fmt.Sprintf("%X%X%X", int(mean.R/257), int(mean.G/257), int(mean.B/257))
	temp, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		panic(err)
	}

	finalMean := uint32(temp)

	// create groups of means
	var nbGroups uint32

	nbGroups = 10000
	finalMean /= (16777215 / nbGroups)

	var size int32

	// create size value
	if height < 500 && width < 500 {
		size = 0
	} else if height > 1000 && width > 1000 {
		size = 2
	} else {
		size = 1
	}

	// get images
	images := getImages(finalMean, size)

	for i := 0; i < len(images); i++ {
		exists := compareImage(pixelsImage, images[i].Path, width, height)
		if exists {
			return true, images[i].Path
		}
	}

	id := getID("image")

	ext := strings.Split(path, ".")

	extension := ext[len(ext)-1]

	if extension == "jpeg" {
		extension = "jpg"
	}

	newName := fmt.Sprintf("./public/%d.%s", id, extension)

	addExistingImage(Image{Path: newName, Color: finalMean, Size: size, ID: id})

	// if gets here, its false
	return false, newName
}

func compareImage(pixelsImage1 [][]Pixel, path2 string, width1 int, height1 int) bool {

	// Get the pixel arrays of the two images
	pixelsImage2, _ := getPixels(path2)

	width2 := len(pixelsImage2[0])
	height2 := len(pixelsImage2)

	// Check if the dimension is equal or not
	if width1 != width2 && height1 != height2 {
		return false
	}

	// The two images are the same size
	var nbPixelsEquivalent int
	var counter int

	// Size of the square pixel groups of which the mean will be computed (can be tweaked)
	pixelSize := 10

	// for each group of pixel, check if it is equivalent to its counterpart of the other picture
	for i := 0; i < height1-height1%pixelSize; i += pixelSize {
		for j := 0; j < width1-width1%pixelSize; j += pixelSize {
			counter++
			mean1 := computeMean(pixelSize, i, j, pixelsImage1)
			mean2 := computeMean(pixelSize, i, j, pixelsImage2)
			if areTheSamePixels(mean1.R, mean1.G, mean1.B, mean2.R, mean2.G, mean2.B) {
				nbPixelsEquivalent++
			}
		}
	}

	result := float32(nbPixelsEquivalent) / float32(counter) * 10000
	// fmt.Println("The two images have a resemblance of", (result)/100, "%")

	if result > 0.91 {
		return true
	}
	return false
}

// Get the bi-dimensional pixel array
func getPixels(filePath string) ([][]Pixel, error) {

	file, _ := os.Open(filePath)
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := make([][]Pixel, height / 2)

	// for each pixel, add it to the array
	for y := 0; y < height / 2; y++ {
		row := make([]Pixel, width /2)
		for x := 0; x < width / 2; x++ {
			row[x] = rgbaToPixel(img.At(x * 2, y * 2).RGBA())
		}
		pixels[y] = row
	}
	return pixels, nil
}

// img.At(x, y).RGBA() returns four int values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r), int(g), int(b), int(a)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func areTheSamePixels(r1 int, g1 int, b1 int, r2 int, g2 int, b2 int) bool {
	if abs(r1-r2) < 10 && abs(g1-g2) < 10 && abs(b1-b2) < 10 {
		return true
	}
	return false

}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func computeMean(side int, x int, y int, image [][]Pixel) Pixel {
	var sumR int = 0
	var sumG int = 0
	var sumB int = 0
	var sumA int = 0

	var count int = 1

	xWidth := x + side
	yHeight := y + side

	// Avoid out of bounds
	if xWidth > len(image) {
		xWidth = len(image) - 1
	}
	if yHeight > len(image[0]) {
		yHeight = len(image[0]) - 1
	}

	// compute the mean
	for i := x; i < xWidth; i++ {
		for j := 0; j < yHeight; j++ {
			sumR += image[i][j].R
			sumG += image[i][j].G
			sumB += image[i][j].B
			sumA += image[i][j].A
			count++
		}
	}

	return Pixel{sumR / count, sumB / count, sumG / count, sumG / count}
}

func computeMeanSlow(width int, height int, x int, y int, image [][]Pixel) Pixel {

	var sumR int = 0
	var sumG int = 0
	var sumB int = 0
	var sumA int = 0

	xWidth := x + width
	yHeight := y + height

	if xWidth > len(image) {
		xWidth = len(image) - 1
	}

	if yHeight > len(image[0]) {
		yHeight = len(image[0]) - 1
	}

	for i := x; i < xWidth; i++ {
		for j := y; j < yHeight; j++ {
			sumR += image[i][j].R
			sumG += image[i][j].G
			sumB += image[i][j].B
			sumA += image[i][j].A
		}
	}

	return Pixel{sumR / int(width*height), sumG / int(width*height), sumB / int(width*height), sumA / int(width*height)}
}
