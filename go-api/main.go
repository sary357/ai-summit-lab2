package main

import (
	"go-api/config"
	routes "go-api/routes"
	"go-api/utils"
	"strconv"
        _ "go-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title go API
// @version 1.0.0
// @description a sample for platform engineers to start with gin framework.

// @contact.name Fu-Ming Tsai
// @contact.email sary357@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @schemes http
func main() {

	// init log
	gin.DisableConsoleColor()
	r := gin.Default()

	// TODO: please put your routes in routes/ folder.

	// for healthcheck
	routes.SetupHealthCheckRoute(r)

	// for API doc
	//routes.SetupSwagRouter(r)
        routes.SetupAwsCdkRoute(r)

        utils.LogInstance.WithFields(logrus.Fields{
		"Host": utils.GetHostname(),
	}).Info("go-api is starting. If there is no error message, it means the service is ready.")

	r.Run("0.0.0.0:" + strconv.Itoa(config.Port)) // listen and serve on 0.0.0.0:8080

}
