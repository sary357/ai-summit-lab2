package utils

import (
        "testing"
	"time"
	"strconv"
	"io/ioutil"
        "os"
)

func TestGenerateUniqueID(t *testing.T) {
        id1 := GenerateRandomFolderId()
        id2 := GenerateRandomFolderId()
        // Basic test: IDs should be different
        if id1 == id2 {
                t.Errorf("Generated IDs are identical: %s", id1)
        }

        // Check format: Should match "app-YYYYMMDDHHmmSS-00000000"
        if len(id1) != 27 {
                t.Errorf("Invalid ID length: %s", id1)
        }

        // Check timestamp portion
        timestamp := id1[4:18]
        _, err := time.Parse("20060102150405", timestamp)
        if err != nil {
                t.Errorf("Invalid timestamp format: %s", timestamp)
        }

        // Check random number portion
        randomNumber, err := strconv.Atoi(id1[19:])
        if err != nil || randomNumber < 0 || randomNumber >= 100000000 {
                t.Errorf("Invalid random number format: %s", id1)
        }
}

func TestSaveContentToFile(t *testing.T) {
        tempDir, err := ioutil.TempDir("", "test")
        if err != nil {
                t.Fatal(err)
        }
        defer os.RemoveAll(tempDir)

        filePath := tempDir + "/test.txt"
        content := ("Hello, world!")

        if !SaveFile(filePath, content) {
                t.Error("Failed to save content to file")
        }

        // Verify the content of the file
        fileBytes, err := ioutil.ReadFile(filePath)
        if err != nil {
                t.Error("Failed to read file")
        }

        if string(fileBytes) != string(content) {
                t.Errorf("File content mismatch: expected %s, got %s", string(content), string(fileBytes))
        }
}
