package route

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api/app"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func prepare() (r *gin.Engine) {

	gin.DisableConsoleColor()

	return gin.Default()
}

func TestSetupHealthCheckRoute(t *testing.T) {
	r := prepare()
	var returnBody app.StatsResponse

	// test health check, please modify it according to your requirement
	SetupHealthCheckRoute(r)
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	json.Unmarshal(w.Body.Bytes(), &returnBody)
	assert.Equal(t, "OK", returnBody.Status)
}
