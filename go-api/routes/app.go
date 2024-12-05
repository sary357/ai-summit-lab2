package route

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-api/app"
	"go-api/utils"
)

type CodeAndRelatedObject struct {
    Code string `json:"code"`
    RequirementTxt string `json:"requirementTxt"`
}


// @Summary save AWS lambda codes and requirements.txt from langchain side
// @Description this return status after checking the system. However, it always returns "OK" at this moment.
// @Tags AWS
// @Version 1.0
// @produce text/plain
// @Success 200 "The API endpoint URL"  string
// @Failure 500 "Internal server error. Please contact the administrator"  string
// @Router /v1/genapiendpoint [post]
func SetupAwsCdkRoute(r *gin.Engine) {
	r.POST("/v1/genapiendpoint", func(c *gin.Context) {
		var codeAndRelatedObject CodeAndRelatedObject

		if err := c.ShouldBindJSON(&codeAndRelatedObject); err != nil {
			utils.LogInstance.WithFields(logrus.Fields{
                		"error":  err.Error(),
        		}).Info("go-api receiving the user's inputs.")
                        c.JSON(http.StatusBadRequest, "Invalid JSON")
                        return
                }

		utils.LogInstance.WithFields(logrus.Fields{
                	"codes": codeAndRelatedObject.Code,
                	"requirementstxt": codeAndRelatedObject.RequirementTxt,
        	}).Info("user's inputs.")

		// start to process 
		status:=app.SaveAndExec(codeAndRelatedObject.Code, codeAndRelatedObject.RequirementTxt)

		c.JSON(http.StatusOK, status)
	})
}


