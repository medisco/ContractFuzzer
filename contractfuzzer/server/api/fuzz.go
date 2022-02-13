package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/fuzz"
	"go.uber.org/zap"
)

type FuzzAPI interface {
	Post(c *gin.Context)
}

type DefaultFuzzAPI struct {
	Logger *zap.Logger
	AbiDir string
	OutDir string
}

func (api DefaultFuzzAPI) Init(logger *zap.Logger, abiDir, outDir string) DefaultFuzzAPI {
	api.Logger = logger
	api.AbiDir = abiDir
	api.OutDir = outDir
	return api
}

func (api DefaultFuzzAPI) Post(c *gin.Context) {
	api.Logger.Info("Starting fuzzing routine")
	go fuzz.Start(api.AbiDir, api.OutDir)
	c.AbortWithStatus(200)
}
