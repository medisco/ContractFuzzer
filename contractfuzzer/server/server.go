package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/server/api"
	"go.uber.org/zap"
)

type Server interface {
	Init(abiDir, outDir, addrMapPath, reporter string) DefaultServer
	Run() error
}

type DefaultServer struct {
	Router         *gin.Engine
	FuzzAPIs       api.FuzzAPI
	HackAPIs       api.HackAPI
	InstrumentAPIs api.InstrumentAPI
}

func (r DefaultServer) Init(logger *zap.Logger, abiDir, outDir, addrMapPath, reporter string) DefaultServer {
	r.Router = gin.Default()
	r.FuzzAPIs = new(api.DefaultFuzzAPI).Init(logger, abiDir, outDir)
	r.HackAPIs = new(api.DefaultHackAPI).Init(logger, addrMapPath, reporter)
	r.InstrumentAPIs = new(api.DefaultInstrumentAPI).Init(logger)
	return r
}

func (r DefaultServer) Run() error {
	r.Router.POST("/fuzz", r.FuzzAPIs.Post)
	r.Router.GET("/hack", r.HackAPIs.Post)
	r.Router.POST("/instrument", r.InstrumentAPIs.Post)
	return r.Router.Run()
}
