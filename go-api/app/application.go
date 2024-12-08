package app

import (
	"os"
	"strings"
	//"fmt"
	"os/exec"
	"github.com/sirupsen/logrus"
	"go-api/config"
	"go-api/utils"
)

type StatsResponse struct {
	Status string `json:"status"`
	Info string `json:"info"`
}

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

func SaveAndExec(codesContent string, requirementTxtContent string) string {
        folderId:=utils.GenerateRandomFolderId()
	lambdaCodesPath:=config.LambdaCodesPath
	requirementTxtPath:=config.RequirementsTxtPath
	// execute aws sdk: Part 1
	status:=InitAwsCdkTask(folderId)
	if !status {
		 utils.LogInstance.WithFields(logrus.Fields{
			 "InitAwsCdkTask": status,
			 "folderId": folderId,
		 }).Info("go-api is trying to execute AWS CDK but failed.")
		 return "ERR: 1001" // failed to execute AWS CDK task
	}
        
        // generate folder name with random postfix
	realLambdaCodesPath:=strings.ReplaceAll(lambdaCodesPath, "TEMPLATE", folderId)
	realRequirementTxtPath:=strings.ReplaceAll(requirementTxtPath, "TEMPLATE", folderId)

	status=SaveAwsLambdaCodes(realLambdaCodesPath, codesContent)

	utils.LogInstance.WithFields(logrus.Fields{
		"SaveAwsLambdaCodeStatus": status,
		"path": realLambdaCodesPath,
		"content": codesContent,
	}).Info("go-api is trying to save AWS lambda codes to the path.")

	if !status {
		return "ERR: 2001" // failed to save files
	}

	// we have requirements.txt
	if len(requirementTxtContent) >= 0 {
		requirementSavedStatus :=  SaveRequirementTxt(realRequirementTxtPath, requirementTxtContent)
		utils.LogInstance.WithFields(logrus.Fields{
			"SaveRequirementTxtStatus": requirementSavedStatus,
			"path": realRequirementTxtPath,
			"content": requirementTxtContent,
		}).Info("go-api is trying to save requirement.txt to the path.")

		if !requirementSavedStatus {
			return "ERR: 2001" // failed to save files
		}
	}
        
	// let's deploy...
	cdkDeployEndpoint := ExecAwsCdk(folderId)
        if len(cdkDeployEndpoint) > 0 {
		return cdkDeployEndpoint
	} else {
		return "ERR: 1001"
	}
}

func InitAwsCdkTask(folderId string) bool{
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

	workingDir, err := os.Getwd()
	if err != nil {
		utils.LogInstance.WithFields(logrus.Fields{
			"error": err,
		}).Error("go-api failed to get current working dir.")
		return false
	}
	cdkScriptPath:=workingDir+"/app/scripts/init_cdk_env.sh"
	cmd := exec.Command("bash", cdkScriptPath, cdkBaseFolder)
	output, err := cmd.CombinedOutput()
	if err != nil {
		utils.LogInstance.WithFields(logrus.Fields{
			"cdkScriptPath": cdkScriptPath,
			"error": err,
		}).Error("go-api failed to execute cdk script.")

		return false
	}

	utils.LogInstance.WithFields(logrus.Fields{
		"cdkScriptPath": cdkScriptPath,
		"stdout": output,
	}).Info("go-api executed cdk script successfully.")
        return true
}

func ExecAwsCdk(folderId string) string{
	cdkBaseFolder:=strings.ReplaceAll(config.AwsCdkFolder, "TEMPLATE", folderId)

	workingDir, err := os.Getwd()
	if err != nil {
		utils.LogInstance.WithFields(logrus.Fields{
			"error": err,
		}).Error("go-api failed to get current working dir.")
		return ""
	}

	cdkScriptPath:=workingDir+"/app/scripts/deploy_cdk_app.sh"
	cmd := exec.Command("bash", cdkScriptPath, cdkBaseFolder)
	output, err := cmd.CombinedOutput()
	if err != nil {
		utils.LogInstance.WithFields(logrus.Fields{
			"cdkScriptPath": cdkScriptPath,
			"error": err,
		}).Error("go-api failed to execute cdk script.")

		return ""
	}
	return strings.TrimSpace(string(output))
}
