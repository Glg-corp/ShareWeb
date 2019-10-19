import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
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