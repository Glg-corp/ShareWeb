package main

import (
    "fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strconv"
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
	mean := computeMean(height, width, 0, 0, pixelsImage)

	hex := fmt.Sprintf("%X%X%X", mean.R, mean.G, mean.B)
	temp, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		panic(err)
	}

	finalMean := uint32(temp)

	fmt.Println(finalMean)

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
		exists := compareImage(path, images[i].Path)
		if exists {
			return true, strconv.Itoa(int(images[i].ID))
		}
    }
    
    id := getID("image")
    
    newName := fmt.Sprintf("./public/%d.%s", id, path[len(path)-3:])

    addExistingImage(Image{Path: newName, Color: finalMean, Size: size, ID: id})

	// if gets here, its false
	return false, strconv.Itoa(int(id))
}

func compareImage(path1 string, path2 string) bool {

	// Supported image formats
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	// Get the pixel arrays of the two images
	pixelsImage1, err := getPixels(path1)
	pixelsImage2, err := getPixels(path2)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	// If the two images don't have the same scale
	if float64(len(pixelsImage1)) / float64(len(pixelsImage1[0])) != float64(len(pixelsImage2)) / float64(len(pixelsImage2[0])) {
		return false
	} else {

		// If the two images have the same scale but are not the same size
		if len(pixelsImage1) != len(pixelsImage2) {
			//TODO
			return false
		}
	}

	// The two images are the same size
	var nbPixelsEquivalent int
	var counter int
	width := len(pixelsImage1[0])
	height := len(pixelsImage1)

	// Size of the square pixel groups of which the mean will be computed (can be tweaked)
	pixelSize := 10

	for i := 0; i < height-height%pixelSize; i += pixelSize {
		for j := 0; j < width-width%pixelSize; j += pixelSize {
			counter++
			if areTheSamePixels(computeMean(pixelSize, pixelSize, i, j, pixelsImage1).R, computeMean(pixelSize, pixelSize, i, j, pixelsImage1).G, computeMean(pixelSize, pixelSize, i, j, pixelsImage1).B, computeMean(pixelSize, pixelSize, i, j, pixelsImage2).R, computeMean(pixelSize, pixelSize, i, j, pixelsImage2).G, computeMean(pixelSize, pixelSize, i, j, pixelsImage2).B) {
				nbPixelsEquivalent++
			}
		}
	}

	result := float64(nbPixelsEquivalent) / float64(counter) * 100 * 100
	fmt.Println("The two images have a resemblance of", math.Round(result)/100, "%")


	if result > 0.95 {
		return true
	} else {
		return false
	}
}

// Get the bi-dimensional pixel array
func getPixels(filePath string) ([][]Pixel, error) {

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
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
	} else {
		return false
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func computeMean(width int, height int, x int, y int, image [][]Pixel) Pixel {

	sumR := 0
	sumG := 0
	sumB := 0
	sumA := 0

	for i := x; i < width+x; i++ {
		for j := y; j < height+y; j++ {
			sumR += image[i][j].R
			sumG += image[i][j].G
			sumB += image[i][j].B
			sumA += image[i][j].A
		}
	}

	return Pixel{sumR / (width * height), sumG / (width * height), sumB / (width * height), sumA / (width * height)}
}
