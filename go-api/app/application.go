package app

import (
	"bytes"
	"fmt"
	"go-api/config"
	"go-api/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StatsResponse struct {
	Status string `json:"status"`
}

// TODO: Please define business logic in this folder or create a new folder for it.
// The following codes are just an example.

// CheckSystemStatus is used to check the system status
func CheckSystemStatus() StatsResponse {
	// TODO: Please define how to check the system status
	return StatsResponse{Status: "OK"}
}

func SaveLambdaCodes() {
}

func SaveRequirementTxt() {
}

func SaveAll() {
}

func ExecAwsCdk() {
}
