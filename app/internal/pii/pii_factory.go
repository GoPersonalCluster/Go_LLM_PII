package pii

import (
	"errors"
	"fmt"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/database"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/models"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/pii/strategy"
)

type PiiFactory struct {
	event *consumer.IntegrationEvent
}

func (c *PiiFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {

	entity := &models.Payload{
		EventID: event.ID,
		Payload: event.Payload,
	}

	if err := database.DB.Create(entity).Error; err != nil {
		return nil, fmt.Errorf("failed to insert payload: %w", err)
	}

	switch event.MetaHeader[len(event.MetaHeader)-1].EventName {
	case "DefaultPiiEvent":
		return c.GetDefaultPiiStrategy(event)
	default:
		return nil, c.GetDefaultErrorResponse(event)
	}
}
func (c *PiiFactory) GetDefaultErrorResponse(event *consumer.IntegrationEvent) error {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return errors.New(event.EventName + "event not found")
}

func (c *PiiFactory) GetDefaultPiiStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.PiiStrategy{}
	return strategy.New(event)
}
