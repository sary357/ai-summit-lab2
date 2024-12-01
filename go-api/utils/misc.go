package utils

import (
	"os"
	"fmt"
        "math/rand"
        "time"
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
