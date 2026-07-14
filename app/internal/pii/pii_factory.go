package pii

import (
	"errors"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/pii/strategy"
)

type PiiFactory struct {
	event *consumer.IntegrationEvent
}

func (c *PiiFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	switch event.MetaHeader[len(event.MetaHeader)-1].EventName {
	case "DefaultPiiEvent":
		return c.GetDefaulPtiiStrategy(event)
	default:
		return nil, c.GetDefaultErrorResponse(event)
	}
}
func (c *PiiFactory) GetDefaultErrorResponse(event *consumer.IntegrationEvent) error {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return errors.New(event.EventName + "event not found")
}

func (c *PiiFactory) GetDefaulPtiiStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.PiiStrategy{}
	return strategy.New(event)
}
