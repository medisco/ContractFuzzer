package topic

import (
	"github.com/gongbell/contractfuzzer/bus"
	"github.com/gongbell/contractfuzzer/bus/event"
)

const TASK_FINISH_TOPIC = "task:finish"

type TaskFinishTopic struct {
	eventBus bus.EventBus
}

func (t TaskFinishTopic) Init(eventBus bus.EventBus) {
	t.eventBus = eventBus
}

func (t TaskFinishTopic) Publish(e event.TaskFinishEvent) {
	t.eventBus.Publish(TASK_FINISH_TOPIC, e)
}
