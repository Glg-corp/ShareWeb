package main

import (
	"math/rand"
	"log"
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

	found := true

	// generate an id
	var id int32
	for found {
		id = int32(rand.Intn(1000000-500000) + 500000)
		found = doesImageExist(id)
		log.Println(id)
	}
	
	// save the image with that id
	image.ID = id
	db.Save(&image)
	return id
}
