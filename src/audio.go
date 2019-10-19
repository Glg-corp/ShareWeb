package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/youpy/go-wav"
)

func compareSounds(path1 string, path2 string) bool {
	// Parse first file
	file1, _ := os.Open(path1)
	reader1 := wav.NewReader(file1)

	// Parse second file
	file2, _ := os.Open(path2)
	reader2 := wav.NewReader(file2)

	// Close proprement les fichiers
	defer file1.Close()
	defer file2.Close()

	// Let's parse ce truc
	samples1, err := reader1.ReadSamples()
	if err == io.EOF {
		return false
	}
	log.Println(len(samples1))

	// L'autre truc aussi
	samples2, err := reader2.ReadSamples()
	if err == io.EOF {
		return false
	}

	// On check si les deux fichiers ont la même longueur
	if !(abs(len(samples1)-len(samples2)) < 32) {
		// La différence est trop grande
		fmt.Println("Length of sounds are too different.")
		return false
	}

	for i := 0; i < myMin(len(samples1), len(samples2)); i++ {
		if cleverCompare(samples1[i].Values[0], samples2[i].Values[0], 32) && cleverCompare(samples1[i].Values[1], samples2[i].Values[1], 32) {

		} else {
			fmt.Println("Sample quelque peu wtf", samples1[i], "et", samples2[i])
			return false
		}
	}

	return true
}

func cleverCompare(a int, b int, marge int) bool {
	return (abs(a-b) < marge)
}

func myMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
