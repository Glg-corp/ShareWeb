package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func routeAddImage(c *gin.Context) {

	// Variable init
	JSONs := []gin.H{{"nom de votre image": "URL d'accès"}}
	// Get data list
	form, _ := c.MultipartForm()
	files := form.File["CONTENT"]

	//Treat each file
	for _, file := range files {

		// Store it locally
		err := c.SaveUploadedFile(file, "temp/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		// Define what type of file it is
		contentType := http.DetectContentType(fileheaderToBytes(file))
		log.Println(contentType)

		isNew := false
		_ = isNew
		idMedia := "caca.jpg"
		_ = idMedia
		// Si c'est un son, on balance
		if contentType == "audio/wave" {
			isNew, idMedia = startCompareSound("temp/" + file.Filename)
		} else if contentType == "image/png" || contentType == "image/jpeg" {
			// Guillaume intervient
		} else {
			log.Println("Holy cucumber... What the fucker ?")
		}

		// Store data
		/// Delete temp file
		err = os.Remove("temp/" + file.Filename)
		if err != nil {
			log.Println("Could not delete temp file : " + file.Filename)
		}
		/// Save file
		if isNew {
			println("Okay les gus, nouveau fichier")
			c.SaveUploadedFile(file, "public/"+idMedia)
		}

		// Make JSON
		JSONs = append(JSONs, gin.H{file.Filename: "caca"})
	}

	// Return Links
	c.JSON(http.StatusAccepted, JSONs)
}

// Convert a file header into an array of byte
func fileheaderToBytes(file *multipart.FileHeader) (returnValue []byte) {
	// Convert FieHeader to file
	fileRightType, err := file.Open()
	if err != nil {
		panic(err)
	}
	{
		buffer := bytes.NewBuffer(nil)
		// Copy file into a buffer
		io.Copy(buffer, fileRightType)
		returnValue = buffer.Bytes()
	}
	return
}
