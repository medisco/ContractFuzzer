package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gongbell/contractfuzzer/pkg/event"
	"github.com/gongbell/contractfuzzer/server/api"
	"go.uber.org/zap"
)

type Server interface {
	Run() error
}

type DefaultServer struct {
	Router         *gin.Engine
	FuzzAPIs       api.FuzzAPI
	HackAPIs       api.HackAPI
	InstrumentAPIs api.InstrumentAPI
}

func (r DefaultServer) Init(logger *zap.Logger, addrMapPath, reporter string, bus event.EventBus) DefaultServer {
	r.Router = gin.Default()
	r.FuzzAPIs = new(api.DefaultFuzzAPI).Init(logger, bus)
	r.HackAPIs = new(api.DefaultHackAPI).Init(logger, addrMapPath, reporter, nil, nil)
	r.InstrumentAPIs = new(api.DefaultInstrumentAPI).Init(logger)
	return r
}

func (r DefaultServer) Run() error {
	r.Router.POST("/fuzz/start", r.FuzzAPIs.Start)
	r.Router.GET("/hack", r.HackAPIs.Post)
	r.Router.POST("/instrument", r.InstrumentAPIs.Post)
	return r.Router.Run()
}
