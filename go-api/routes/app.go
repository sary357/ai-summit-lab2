package route

import (
	"fmt"
	"encoding/json"
	"net/http"
	//"go-api/app"
	"go-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&codeAndRelatedObject); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, "Invalid JSON")
			return 
		}

                fmt.Println("\nCodes: %s\nrequirements.txt: %s\n", codeAndRelatedObject.Code, codeAndRelatedObject.RequirementTxt)
		utils.LogInstance.WithFields(logrus.Fields{
                	"codes": codeAndRelatedObject.Code,
                	"requirementstxt": codeAndRelatedObject.RequirementTxt,
        	}).Info("go-api receiving the user's inputs.")

		c.JSON(200, "https://xxx.xxx.xxx/abc")
	})
}


