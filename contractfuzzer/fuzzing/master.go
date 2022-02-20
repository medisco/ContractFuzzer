package fuzzing

import (
	"fmt"
	"time"

	"github.com/gongbell/contractfuzzer/fuzz"
	"github.com/gongbell/contractfuzzer/pkg/event"
	"go.uber.org/zap"
)

type FuzzingMaster interface {
}

type DefaultFuzzingMaster struct {
	Logger   *zap.Logger
	EventBus event.EventBus
	AbiDir   string
	OutDir   string
}

func (m DefaultFuzzingMaster) Init(
	logger *zap.Logger,
	eventBus event.EventBus,
	abiDir string,
	outDir string,
) FuzzingMaster {
	m.Logger = logger
	m.EventBus = eventBus
	m.AbiDir = abiDir
	m.OutDir = outDir

	m.EventBus.Subscribe("task:request", m.startFuzzer)
	return m
}

func (m DefaultFuzzingMaster) startFuzzer(taskId string, contracts []string, duration time.Duration) {
	m.Logger.Info(fmt.Sprintf("Running fuzzing task %s for %-8v", taskId, duration))
	m.EventBus.SubscribeOnce(fmt.Sprintf("task:finish:%s", taskId), func() {
		m.Logger.Info(fmt.Sprintf("Stopping fuzzing task %s", taskId))
		fuzz.G_stop <- true
	})
	go m.startTimer(taskId, duration)
	fuzz.Start(m.AbiDir, m.OutDir)
}

func (m DefaultFuzzingMaster) startTimer(taskId string, duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
	m.EventBus.Publish(fmt.Sprintf("task:finish:%s", taskId))
	m.EventBus.Publish("task:finish", taskId)
}
