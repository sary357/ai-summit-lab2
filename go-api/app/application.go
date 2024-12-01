package app

import (
//	"bytes"
//	"fmt"
	"go-api/config"
	"go-api/utils"
//	"io"
//	"net/http"
        "strings"
//	"github.com/gin-gonic/gin"
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

func SaveAwsLambdaCodes(path string, content string) bool{
	status := utils.SaveFile(path, content)
	return status
}

func SaveRequirementTxt(path string, content string) bool{
	status := utils.SaveFile(path, content)
	return status
}

func SaveAll (codesContent string, requirementTxtContent string) bool {
        folderId:=utils.GenerateRandomFolderId()
	lambdaCodesPath:=config.LambdaCodesPath
	requirementTxtPath:=config.RequirementsTxtPath
	// TODO: execute aws sdk

        // generate folder name with random postfix
	realLambdaCodesPath:=strings.ReplaceAll(lambdaCodesPath, "TEMPLATE", folderId)
	realRequirementTxtPath:=strings.ReplaceAll(requirementTxtPath, "TEMPLATE", folderId)

	status:=SaveAwsLambdaCodes(realLambdaCodesPath, codesContent)

	utils.LogInstance.WithFields(logrus.Fields{
		"SaveAwsLambdaCode status": status,
		"path": realLambdaCodesPath,
		"content": codesContent,
	}).Info("go-api is trying to save AWS lambda codes to the path.")

	if status {
		if len(requirementTxtContent) > 0 {
			requirementSavedStatus :=  SaveRequirementTxt(realRequirementTxtPath, requirementTxtContent)
			utils.LogInstance.WithFields(logrus.Fields{
				"SaveRequirementTxt": requirementSavedStatus,
				"path": realRequirementTxtPath,
				"content": requirementTxtContent,
			}).Info("go-api is trying to save requirement.txt to the path.")

			return requirementSavedStatus
		} else {
			return true
		}
	} else {
		return false
	}
}

//func ExecAwsCdk() {
//}
