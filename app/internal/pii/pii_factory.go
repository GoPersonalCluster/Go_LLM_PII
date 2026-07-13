package pii

import (
	"errors"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/config"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/log/strategy"
)

type LogFactory struct {
	event *consumer.IntegrationEvent
}

func (c *LogFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	switch event.EventName {
	case "ErrorLog":
		return c.GetErrorQueue(event)
	case "PipelineLog":
		return c.GetPipelineQueue(event)
	default:
		return nil, c.GetDefaultErrorResponse(event)
	}
}
func (c *LogFactory) GetDefaultErrorResponse(event *consumer.IntegrationEvent) error {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return errors.New(event.EventName + "event not found")
}

func (c *LogFactory) GetErrorQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.ErrorLogStrategy{}
	return strategy.New(event)
}

func (c *LogFactory) GetPipelineQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.PipelineLogStrategy{}
	return strategy.New(event)
}
