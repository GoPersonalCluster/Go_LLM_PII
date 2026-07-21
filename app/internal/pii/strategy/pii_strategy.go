package strategy

import (
	"context"
	"fmt"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	llm "github.com/GoPersonalCluster/go_llm_pii/app/internal/client"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/database"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/models"
	"github.com/google/uuid"
)

type PiiStrategy struct {
	event *consumer.IntegrationEvent
}

func (pS *PiiStrategy) New(iE *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {

	mh := iE.CreateMetaHeader(config.GetHostName(), "DefaultPiiEvent")
	iE.MetaHeader = append(iE.MetaHeader, mh)

	return &PiiStrategy{event: iE}, nil
}

func (pS *PiiStrategy) Start() ([]byte, error) {

	client := llm.NewClient()

	response, err := client.Chat(context.Background(), string(pS.event.Payload))
	if err != nil {
		return nil, err
	}
	eventID, err := uuid.Parse(pS.event.ID)
	if err != nil {
		return nil, err
	}

	entity := &models.TagExtraction{
		PayloadID: eventID,
		Result:    response,
	}

	if err := database.DB.Create(entity).Error; err != nil {
		return nil, fmt.Errorf("failed to insert payload: %w", err)
	}

	println("added to db")
	return []byte(response), nil
}
