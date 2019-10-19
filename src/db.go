package main

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// database pointer
var db *gorm.DB

// image table struct
type Image struct {
	ID    int32 `gorm:"primary_key:yes;column:ID"`
	Path  string
	Color uint32
	Size  int32
	Key   string
}

// sound table struct
type Sound struct {
	ID        int32 `gorm:"primary_key:yes;column:ID"`
	Path      string
	NbSamples int32
	Mono      bool
	Key       string
}

// fallback table struct
type Other struct {
	ID        int32 `gorm:"primary_key:yes;column:ID"`
	Path      string
	Extension string
	FileSize  int32
	Key       string
}

// init the database
func initDB() {
	rand.Seed(time.Now().UTC().UnixNano())
	var err error
	db, err = gorm.Open("sqlite3", "salut.db")
	// if there is an error, throw the exception
	if err != nil {
		panic(err)
	}

	// migrate the table
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Sound{})
	db.AutoMigrate(&Other{})
}

func getImage(id int32) Image {

	// get the rows
	rows, err := db.Where(&Image{ID: id}).Model(&Image{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	// return the image
	var image Image
	if rows.Next() {
		db.ScanRows(rows, &image)
	}
	return image
}

// return a bool whether an image exist
func doesImageExist(id int32) bool {
	rows, err := db.Where(&Image{ID: id}).Model(&Image{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return rows.Next()
}

func getImages(color uint32, size int32) []Image {

	// array := []Image{}

	// get the matching rows
	rows, err := db.Where(&Image{Color: color, Size: size}).Model(&Image{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var images []Image

	// add images to the list
	var image Image
	for rows.Next() {
		db.ScanRows(rows, &image)
		images = append(images, image)
	}

	return images

}

// add a new image
func addImage(image Image) int32 {

	id := getID("image")

	// save the image with that id
	image.ID = id
	db.Save(&image)
	return id
}

/// --- sound ---

// addSound
func addSound(sound Sound) int32 {

	id := getID("sound")

	// save the image with that id
	sound.ID = id
	db.Save(&sound)
	return id
}

// return a bool if a image exists
func doesSoundExist(id int32) bool {
	rows, err := db.Where(&Sound{ID: id}).Model(&Sound{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return rows.Next()
}

func getSound(id int32) Sound {
	// get the rows
	rows, err := db.Where(&Sound{ID: id}).Model(&Sound{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	// return the image
	var sound Sound
	if rows.Next() {
		db.ScanRows(rows, &sound)
	}
	return sound
}

func getID(mode string) int32 {
	found := true

	// generate an id
	var id int32
	for found {

		id = int32(rand.Intn(1000000-500000) + 500000)
		if mode == "image" {
			found = doesImageExist(id)
		} else if mode == "sound" {
			found = doesSoundExist(id)
		} else if mode == "other" {
			found = doesOtherExist(id)
		} else {
			panic("invalid mode")
		}
	}
	return id

}

func getSounds(nbSamples int32, mono bool) []Sound {

	// get the matching rows
	rows, err := db.Where(&Sound{NbSamples: nbSamples, Mono: mono}).Model(&Sound{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var sounds []Sound

	// add images to the list
	var sound Sound
	for rows.Next() {
		db.ScanRows(rows, &sound)
		sounds = append(sounds, sound)
	}

	return sounds

}

func addExistingImage(image Image) {
	db.Save(&image)
}

func addExistingSound(sound Sound) {
	db.Save(&sound)
}

// ----- others -------

func doesOtherExist(id int32) bool {
	rows, err := db.Where(&Other{ID: id}).Model(&Other{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return rows.Next()
}

func getOthers(extension string, fileSize int32) []Other {
	// get the matching rows
	rows, err := db.Where(&Other{Extension: extension, FileSize: fileSize}).Model(&Other{}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var others []Other

	// add the lines to the list
	var other Other
	for rows.Next() {
		db.ScanRows(rows, &other)
		others = append(others, other)
	}

	return others
}

func addExistingOther(other Other) {
	db.Save(&other)
}

// Security  functions

func isAllowed(mode string, path string, key string) bool {
	var count int32
	if mode == "image" {
		db.Where(&Image{Path: path, Key: key}).Count(&count)
	} else if mode == "sound" {
		db.Where(&Sound{Path: path, Key: key}).Count(&count)
	} else if mode == "other" {
		db.Where(&Other{Path: path, Key: key}).Count(&count)
	} else {
		panic("Invalid mode >:c")
	}
	return count == 1
}
