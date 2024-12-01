package utils

import (
	"os"
	"fmt"
        "math/rand"
        "time"
	"io/ioutil"
	"path/filepath"
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

        // Generate a random number between 0 and 99999999
        rand.Seed(time.Now().UnixNano())
        randomNumber := rand.Intn(100000000)

        // Concatenate the timestamp and random number
        randomId := fmt.Sprintf("%s-%08d", timestamp, randomNumber)
	return randomId
}

func SaveFile(path string, content string) bool {
	file_content := []byte(content)
        err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
        if err != nil {
                fmt.Println("Error creating directory:", err)
                return false
        }

        // Write the content to the file
        err = ioutil.WriteFile(path, file_content, 0644)
        if err != nil {
                fmt.Println("Error writing to file:", err)
                return false
        }

        return true
}
