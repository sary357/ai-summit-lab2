package route

import (
	"fmt"
	//"go-api/app"
	"go-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


// @Summary save AWS lambda codes and requirements.txt from langchain side
// @Description this return status after checking the system. However, it always returns "OK" at this moment.
// @Tags AWS
// @Version 1.0
// @produce text/plain
// @Success 200 "The API endpoint URL"  string
// @Failure 500 "Internal server error. Please contact the administrator"  string
// @Router /v1/execodes [post]
func SetupAwsCdkRoute(r *gin.Engine) {
	r.POST("/v1/execodes", func(c *gin.Context) {
		codes := c.PostForm("codes")
		requirementTxt := c.PostForm("requirementstxt")
                fmt.Println(codes)
		fmt.Println(requirementTxt)
		utils.LogInstance.WithFields(logrus.Fields{
                	"codes": codes,
                	"requirementstxt": requirementTxt,
        	}).Info("go-api receiving the user's inputs.")

		c.JSON(200, "https://xxx.xxx.xxx/abc")
	})
}


