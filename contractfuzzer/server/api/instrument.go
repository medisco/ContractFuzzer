package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/server/model"
	"go.uber.org/zap"
)

type InstrumentAPI interface {
	Post(c *gin.Context)
}

type DefaultInstrumentAPI struct {
	Logger *zap.Logger
}

func (api DefaultInstrumentAPI) Init(logger *zap.Logger) DefaultInstrumentAPI {
	api.Logger = logger
	return api
}

func (api DefaultInstrumentAPI) Post(c *gin.Context) {
	var request model.InstrumentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	api.Logger.Info(fmt.Sprintf("json: %s", string(request.Name)))
	c.AbortWithStatus(200)
}
