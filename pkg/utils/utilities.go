package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"
)

// create folder if not exist
func CreateFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

// string to hash
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// get timestamp
func TimeStamp() string {
	return (time.Now()).Format("20060102150405")
}
