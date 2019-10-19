package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"mime/multipart"

	"io"

	"bytes"

	"os"

	"strconv"
)

func routeAddImage(c *gin.Context) {

	// Variable init
	JSONs := []gin.H{{"nom de votre image": "URL d'accès"}}
	// Get data list
	form, _ := c.MultipartForm()
	files := form.File["CONTENT"]

	//Treat each file
	for i, file := range files {

		// Store it locally
		err := c.SaveUploadedFile(file, "temp/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		// Define what type of file it is
		contentType := http.DetectContentType(fileheaderToBytes(file))
		log.Println(contentType)

		// Call treatment functions
		url := "https://theuselessweb.com/"
		isPresent := false

		// Store data
		/// Delete temp file
		err = os.Remove("temp/" + file.Filename)
		if err != nil {
			log.Println("Could not delete temp file : " + file.Filename)
		}
		/// Save file
		if !isPresent {
			c.SaveUploadedFile(file, "public/"+strconv.Itoa(i))
		}

		// Make JSON
		JSONs = append(JSONs, gin.H{file.Filename: url})
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
