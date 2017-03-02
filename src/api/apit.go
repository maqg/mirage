package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) LoadApiTestPage(c *gin.Context) {
	apiModules, _ := json.Marshal(GApiConfig.Modules)
	c.HTML(http.StatusOK, "index.html",
		gin.H{
			"TESTTITLE": "Mirage",
			"APICONFIG": string(apiModules),
		})
}
