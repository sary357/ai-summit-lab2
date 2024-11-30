package route

import (
	"go-api/app"

	"github.com/gin-gonic/gin"
)

// StatsResponse is the response for health check

// @Summary health check for external monitoring system
// @Description this return status after checking the system. However, it always returns "OK" at this moment.
// @Tags Monitoring
// @Version 1.0
// @produce text/plain
// @Success 200 {object}  StatsResponse
// @Failure 500 {object}  StatsResponse
// @Router /healthcheck [get]
func SetupHealthCheckRoute(r *gin.Engine) {
	status := app.CheckSystemStatus()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, status)
	})
}
