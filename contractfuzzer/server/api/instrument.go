package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/server/model"
)

type InstrumentAPI interface {
	Post(c *gin.Context)
}

type DefaultInstrumentAPI struct {
}

func (api DefaultInstrumentAPI) Init() DefaultInstrumentAPI {
	return api
}

func (api DefaultInstrumentAPI) Post(c *gin.Context) {
	var request model.InstrumentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println("json: ", string(request.Name))
	c.AbortWithStatus(200)
}
