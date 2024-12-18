package route

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-api/app"
	"go-api/utils"
)

type CodeAndRelatedObject struct {
	Codes string `json:"codes"`
	RequirementTxt string `json:"requirementTxt"`
}

type ReturnObj struct {
	Endpoint string `json:"endpoint"`
	Message string `json:"message"`
}

// @Summary save AWS lambda codes and requirements.txt from langchain side
// @Description this return status after checking the system. However, it always returns "OK" at this moment.
// @Tags AWS
// @Version 1.0
// @produce text/plain
// @Success 200 "The API endpoint URL"  string
// @Failure 500 "Internal server error. Please contact the administrator"  string
// @Router /v1/genapi [post]
func SetupAwsCdkRoute(r *gin.Engine) {
	r.POST("/v1/genapi", func(c *gin.Context) {
		var codeAndRelatedObject CodeAndRelatedObject

		if err := c.ShouldBindJSON(&codeAndRelatedObject); err != nil {
			utils.LogInstance.WithFields(logrus.Fields{
                		"error":  err.Error(),
        		}).Info("go-api receiving the user's inputs.")
                        c.JSON(http.StatusBadRequest, "Invalid JSON")
                        return
                }

		utils.LogInstance.WithFields(logrus.Fields{
                	"codes": codeAndRelatedObject.Codes,
                	"requirementstxt": codeAndRelatedObject.RequirementTxt,
        	}).Info("user's inputs.")

		// start to process 
		status:=app.SaveAndExec(codeAndRelatedObject.Codes, codeAndRelatedObject.RequirementTxt)
		//status:= "https://ihznxmqgj9.execute-api.ap-northeast-1.amazonaws.com/prod/"
		if strings.Contains(status, "ERR") {
			retObj := ReturnObj{
				Endpoint: "",
				Message: status,
			}
                        c.JSON(http.StatusBadRequest, retObj)
		} else {
			retObj := ReturnObj{
                                Endpoint: status,
                                Message: "",
                        }
			c.JSON(http.StatusOK, retObj)
		}
	})
}


