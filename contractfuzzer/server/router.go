package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/server/api"
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

func (r DefaultServer) Init(abiDir, outDir, addrMapPath, reporter string) DefaultServer {
	r.Router = gin.Default()
	r.FuzzAPIs = new(api.DefaultFuzzAPI).Init(abiDir, outDir)
	r.HackAPIs = new(api.DefaultHackAPI).Init(addrMapPath, reporter)
	r.InstrumentAPIs = new(api.DefaultInstrumentAPI).Init()
	return r
}

func (r DefaultServer) Run() error {
	r.Router.POST("/fuzz", r.FuzzAPIs.Post)
	r.Router.POST("/hack", r.HackAPIs.Post)
	r.Router.POST("/instrument", r.InstrumentAPIs.Post)
	return r.Router.Run()
}
