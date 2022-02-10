package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/fuzz"
)

type FuzzAPI interface {
	Post(c *gin.Context)
}

type DefaultFuzzAPI struct {
	AbiDir string
	OutDir string
}

func (api DefaultFuzzAPI) Init(abiDir, outDir string) DefaultFuzzAPI {
	api.AbiDir = abiDir
	api.OutDir = outDir
	return api
}

func (api DefaultFuzzAPI) Post(c *gin.Context) {
	log.Printf("Starting go routine")
	go fuzz.Start(api.AbiDir, api.OutDir)
	c.AbortWithStatus(200)
}
