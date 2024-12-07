package app

import (
//	"bytes"
//	"fmt"
	"os"
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
	Info string `json:"info"`
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

func SaveAndExec(codesContent string, requirementTxtContent string) bool {
        folderId:=utils.GenerateRandomFolderId()
	lambdaCodesPath:=config.LambdaCodesPath
	requirementTxtPath:=config.RequirementsTxtPath
	// TODO: execute aws sdk: Part 1
	
        
        // generate folder name with random postfix
	realLambdaCodesPath:=strings.ReplaceAll(lambdaCodesPath, "TEMPLATE", folderId)
	realRequirementTxtPath:=strings.ReplaceAll(requirementTxtPath, "TEMPLATE", folderId)

	status:=SaveAwsLambdaCodes(realLambdaCodesPath, codesContent)

	utils.LogInstance.WithFields(logrus.Fields{
		"SaveAwsLambdaCodeStatus": status,
		"path": realLambdaCodesPath,
		"content": codesContent,
	}).Info("go-api is trying to save AWS lambda codes to the path.")

	if status {
		if len(requirementTxtContent) > 0 {
			requirementSavedStatus :=  SaveRequirementTxt(realRequirementTxtPath, requirementTxtContent)
			utils.LogInstance.WithFields(logrus.Fields{
				"SaveRequirementTxtStatus": requirementSavedStatus,
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

func ExecAwsCdkTask(codePath string, folderId string) bool{
	cdkBaseFolder:=strings.ReplaceAll(config.AwsCdkFolder, "TEMPLATE", folderId)
	utils.LogInstance.WithFields(logrus.Fields{
		"path": cdkBaseFolder,
	}).Info("go-api is creating directory.")
	err := os.MkdirAll(cdkBaseFolder, os.ModePerm)
	if err != nil {
		utils.LogInstance.WithFields(logrus.Fields{
			"path": cdkBaseFolder,
			"error": err,
		}).Error("go-api failed to create directories.")
                return false
        }
        return true
        

}
//func ExecAwsCdk() {
//}
