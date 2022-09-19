package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type KafkaMessage interface {
	InitKafka(req KafkaRequest) error
}
type IRepo interface{}
type kafkaMessage struct {
	consumerMess *KafkaComsumerGroup
	repo         IRepo
}

func NewKafka(
	repo IRepo,
) KafkaMessage {
	return &kafkaMessage{
		repo: repo,
	}
}

func (a *kafkaMessage) InitKafka(req KafkaRequest) error {
	consumerGroup, err := NewKafkaComsumerGroup(context.Background(), req)
	if err != nil {
		zap.S().Errorw(fmt.Sprintf("Error while init connection to brokers: %v", zap.Error(err)))
		return err
	}
	zap.S().Info("Init consumergroup successfully")
	a.consumerMess = consumerGroup
	go a.messageFrom()
	a.consumerMess.CloseConsumerGroup()
	return nil
}

func (a *kafkaMessage) messageFrom() {
	for {
		message, ok := <-a.consumerMess.MessageCh
		if message != nil && ok {
			switch message.Topic {
			case "logging":
				zap.S().Info("consumer message topic logging")

				log := LoggingDTO{}
				err := json.Unmarshal(message.Value, &log)
				if err != nil {
					// a.repo.NewLoggingRepo().AddLogging(context.Background(), &models.Logging{
					// 	PartnerCode: log.PartnerCode,
					// 	Method:      log.Method,
					// 	Data:        string(message.Value),
					// 	Endpoint:    log.Endpoint,
					// 	CreatedAt:   log.CreatedAt,
					// 	UpdatedAt:   log.UpdatedAt,
					// })
					return
				}
				// a.repo.NewLoggingRepo().AddLogging(context.Background(), &models.Logging{
				// 	PartnerCode: log.PartnerCode,
				// 	Method:      log.Method,
				// 	Data:        log.Data,
				// 	Endpoint:    log.Endpoint,
				// 	CreatedAt:   log.CreatedAt,
				// 	UpdatedAt:   log.UpdatedAt,
				// })
			default:
				zap.S().Infof("topic not found")
			}
		}
	}
}
