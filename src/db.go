package main

import (
	"log"
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
	Color string
	Size  int32
}

// sound table struct
type Sound struct {
	ID        int32 `gorm:"primary_key:yes;column:ID"`
	Path      string
	NbSamples int32
	Mono      bool
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

	// get sample values
	populateDb()

}

// generate sample values for debugging purpose
func populateDb() {
	cat := Image{ID: 34567, Path: "./public/cat.jpg", Color: "0000FF", Size: 3}
	db.Save(&cat)

	cat = Image{ID: 57821, Path: "./public/catze.jpg", Color: "0000FF", Size: 3}
	db.Save(&cat)

	cat = Image{ID: 19721, Path: "./public/catt.jpg", Color: "000eFF", Size: 3}
	db.Save(&cat)
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

func getImages(color string, size int32) []Image {

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

func getID(mode string) {
	found := true

	// generate an id
	var id int32
	for found {

		id = int32(rand.Intn(1000000-500000) + 500000)
		if mode == "image" {
			found = doesImageExist(id)
		}
		else if mode == "sound"{
			found = doesSoundExist(id)
		}
		else{
			panic("invalid mode")
		}
		log.Println(id)
	}

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

addExistingImage(image Image){
	db.Save(&Image)
}

addExistingSound(sound Sound){
	db.Save(&Sound)
}