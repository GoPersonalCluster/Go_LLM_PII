package strategy

import (
	"context"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	llm "github.com/GoPersonalCluster/go_llm_pii/app/internal/client"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
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
	prompt := string(pS.event.Payload)
	response, err := client.Chat(context.Background(), prompt)
	if err != nil {
		return nil, err
	}

	return []byte(response), nil
}
