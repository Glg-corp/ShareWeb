package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
)

// take the path of a file in parameter and return his MD5 checksum
func hash_file_md5(filePath string) (string, error) {
	// Initialize ou future return value
	var returnMD5String string
	// Check for errors,if err is not null return the error
	file, err := os.Open(filePath)
	// Once the function is executed the file get closed
	defer file.Close()
	if err != nil {
		return returnMD5String, err
	}

	// Create a new md5 hash in the variable
	hash := md5.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	// get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// convert the bytes data representation to a string representation
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

// take the filepath in paramters and return the size in byte
func file_size_calc(filepath) int32 {
	filepath, err := os.Open(filepath)
	if err != nil {
		// handle the error here
		return
	}
	defer filepath.Close()
	// get the file size
	stat, err := filepath.Stat()
	if err != nil {
		return
	}

	return (stat.Size())

}

func checksum_compare(file_path_to_add, extension) (bool, string) {

	// taille et extension du file path to add
	file_size := file_size_calc(file_path_to_add)

	// create a board
	others := getOthers(file_extension, file_size)

	hash := hash_file_md5(file_path_to_add)

	for i := 0; i < len(others); i++ {
		if hash == hash_file_md5(others[i].Path) {
			return true, strconv.Itoa(int(others[i].ID))
		}
	}

	id := getID("other")

	newName := fmt.Sprintf("./public/%d.%s", id, extension)

	addExistingOther(Other{Path: newName, FileSize: file_size, Extension: extension, ID: id})

	// if gets here, its false
	return false, strconv.Itoa(int(id))

}
