package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// database pointer
var db *gorm.DB

// image table struct
type Image struct {
	ID    int32    `gorm:"primary_key:yes;column:ID"`
	Path  string
	Color string
	Size  int
}

// init the database
func initDB() {
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
}

func getImages(color string, size int32) []string {
	return []string{"salut", "coucou"}
}

func addImage(image Image) {

}
