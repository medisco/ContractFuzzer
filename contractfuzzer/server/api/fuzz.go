package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/pkg/event"
	"github.com/gongbell/contractfuzzer/server/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type FuzzAPI interface {
	Start(c *gin.Context)
}

type DefaultFuzzAPI struct {
	Logger   *zap.Logger
	EventBus event.EventBus
}

func (api DefaultFuzzAPI) Init(logger *zap.Logger, eventBus event.EventBus) DefaultFuzzAPI {
	api.Logger = logger
	api.EventBus = eventBus
	return api
}

func (api DefaultFuzzAPI) Start(c *gin.Context) {
	var request model.FuzzStartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskId := uuid.NewString()
	api.EventBus.Publish("task:request", taskId, request.Contracts, request.Duration)
	c.JSON(200, model.FuzzStartResponse{TaskId: taskId})
}
