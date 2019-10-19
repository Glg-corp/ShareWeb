package main

import (
	"fmt"
	"io"
	"os"

	"github.com/youpy/go-wav"
)

func startCompareSound(path string) (bool, string) {
	// Open file
	file, _ := os.Open(path)
	reader := wav.NewReader(file)
	defer file.Close()

	// On the parsing
	samples, _ := reader.ReadSamples()
	nbSamples := len(samples)
	isStereo := isSoundStereo(samples)

	sounds := getSounds(int32(nbSamples), !isStereo)

	// Let's compare all sound with dat sound
	for _, sound := range sounds {
		if compareSounds(samples, sound.Path) {
			// On est bon, on l'a trouvé
			return true, sound.Path
		}
	}

	// Let's build a new son
	sound := Sound{
		Mono:      !isStereo,
		NbSamples: int32(nbSamples),
		Path:      path}

	addSound(sound)

	return false, path
}

func compareSounds(samples1 []wav.Sample, path2 string) bool {
	// Parse second file
	file2, _ := os.Open(path2)
	reader2 := wav.NewReader(file2)

	// Close proprement les fichiers
	defer file2.Close()

	// Let's parse ce truc
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

func isSoundStereo(samples []wav.Sample) bool {
	for _, sample := range samples {
		if sample.Values[0] != sample.Values[1] {
			return true
		}
	}
	return false
}
