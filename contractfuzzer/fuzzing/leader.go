package fuzzing

import (
	"fmt"
	"time"

	"github.com/gongbell/contractfuzzer/bus"
	"github.com/gongbell/contractfuzzer/fuzz"
	"go.uber.org/zap"
)

type FuzzingLeader interface {
}

type DefaultFuzzingLeader struct {
	Logger   *zap.Logger
	EventBus bus.EventBus
	AbiDir   string
	OutDir   string
}

func (m DefaultFuzzingLeader) Init(
	logger *zap.Logger,
	eventBus bus.EventBus,
	abiDir string,
	outDir string,
) FuzzingLeader {
	m.Logger = logger
	m.EventBus = eventBus
	m.AbiDir = abiDir
	m.OutDir = outDir

	m.EventBus.Subscribe("task:request", m.startFuzzer)
	return m
}

func (m DefaultFuzzingLeader) startFuzzer(taskId string, contracts []string, duration time.Duration) {
	m.Logger.Info(fmt.Sprintf("Running fuzzing task %s for %-8v", taskId, duration))
	m.EventBus.SubscribeOnce(fmt.Sprintf("task:finish:%s", taskId), func() {
		m.Logger.Info(fmt.Sprintf("Stopping fuzzing task %s", taskId))
		fuzz.G_stop <- true
	})
	go m.startTimer(taskId, duration)
	fuzz.Start(m.AbiDir, m.OutDir, taskId)
}

func (m DefaultFuzzingLeader) startTimer(taskId string, duration time.Duration) {
	m.Logger.Info(fmt.Sprintf("Start timer for %-8v", duration))

	cancel := make(chan bool)
	m.EventBus.SubscribeOnce(fmt.Sprintf("task:finish:%s", taskId), func() {
		cancel <- true
	})

	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			m.Logger.Info(fmt.Sprintf("Stopping fuzzer task %s", taskId))
			m.EventBus.Publish(fmt.Sprintf("task:finish:%s", taskId))
			m.EventBus.Publish("task:finish", taskId)
			m.Logger.Info(fmt.Sprintf("Stopping fuzzing timer %s", taskId))
			break
		case <-cancel:
			m.Logger.Info(fmt.Sprintf("Stopping fuzzing timer %s", taskId))
			break
		}
	}
}
