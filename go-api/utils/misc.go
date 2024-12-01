package utils

import (
	"os"
	"fmt"
        "math/rand"
        "time"
	"io/ioutil"
	"path/filepath"
	"github.com/sirupsen/logrus"
)

func GetHostname() (hostName string) {
	hostName, err := os.Hostname()
	if err != nil {
		LogInstance.Error("Failed to get host name")
		hostName = ""
	}
	return
}

func GenerateRandomFolderId()  (string) {
        // Get current timestamp
        now := time.Now()
        timestamp := now.Format("app-20060102150405")

        // Generate a random number between 0 and 999
        rand.Seed(time.Now().UnixNano())
        randomNumber := rand.Intn(1000)

        // Concatenate the timestamp and random number
        randomId := fmt.Sprintf("%s-%03d", timestamp, randomNumber)
	return randomId
}

func SaveFile(path string, content string) bool {
	file_content := []byte(content)
        err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
        if err != nil {
		LogInstance.WithFields(logrus.Fields{
			"path": filepath.Dir(path),
			"error": err,
		}).Error("go-api failed to create directories.")
                //fmt.Println("Error creating directory:", err)
                return false
        }

        // Write the content to the file
        err = ioutil.WriteFile(path, file_content, 0644)
        if err != nil {
		LogInstance.WithFields(logrus.Fields{
			"path": path,
			"error": err,
		}).Error("go-api failed to save file")
                return false
        }

	LogInstance.WithFields(logrus.Fields{
			"path": path,
	}).Debug("go-api saved file")
        return true
}
