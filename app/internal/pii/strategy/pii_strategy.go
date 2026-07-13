package strategy

import (
	"encoding/json"
	"fmt"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/config"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/database"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/models"
)

type PiiStrategy struct {
	event *consumer.IntegrationEvent
}

func (pQS *PiiStrategy) New(iE *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	iE.EventName = "PII"
	mh := iE.CreateMetaHeader(config.GetHostName(), "ErrorLogEvent")
	iE.MetaHeader = append(iE.MetaHeader, mh)

	return &PiiStrategy{event: iE}, nil
}

func (pQS *PiiStrategy) Start() ([]byte, error) {

	json, err := json.Marshal(pQS.event.MetaHeader)
	if err != nil {

		last := pQS.event.MetaHeader[len(pQS.event.MetaHeader)-2]
		log := models.ErrorLog{
			Event:       pQS.event.EventName,
			Description: fmt.Sprintf("error during parsing:%s %s %d", last.EventName, last.Source, last.OccuredAt),
		}
		database.DB.Create(&log)

		return nil, err
	} else {

		log := models.ErrorLog{
			Event:       pQS.event.EventName,
			Description: string(json),
		}
		database.DB.Create(&log)

	}

	return nil, nil
}
