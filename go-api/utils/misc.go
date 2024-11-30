package utils

import (
	"os"
)

func GetHostname() (hostName string) {
	hostName, err := os.Hostname()
	if err != nil {
		LogInstance.Error("Failed to get host name")
		hostName = ""
	}
	return
}
