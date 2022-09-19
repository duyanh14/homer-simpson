package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/Shopify/sarama.v1"
)

type LoggingDTO struct{}

type KafkaProducer interface {
	Send(ctx context.Context, req KafkaPropertySend, property LoggingDTO) error
}

type KafkaPropertySend struct {
	ProducerName string
	Topic        string
	Key          string
}

type asyncKafkaProducer struct {
	instanceProducer sarama.AsyncProducer
	topics           []string
	producerName     string
	ready            bool
}

type syncKafkaProducer struct {
	instanceProducer sarama.SyncProducer
	topics           []string
	producerName     string
	ready            bool
}

type KafkaRequest struct {
	Async    bool
	Topics   []string
	Broker   []string
	Version  string
	Username string
	Password string
	Group    string
}

func NewKafkaProducer(req KafkaRequest) (KafkaProducer, error) {
	if req.Async {
		return initAsyncKafkaProducer(req)
	}
	return initSyncKafkaProducer(req)

}

func defaultVersion(v string) string {
	if strings.TrimSpace(v) == "" {
		return "1.34.0" // release version date: 2022-05-30
	}
	return v
}

func initAsyncKafkaProducer(req KafkaRequest) (*asyncKafkaProducer, error) {
	kafkaPro := &asyncKafkaProducer{}
	if len(req.Broker) == 0 {
		return kafkaPro, fmt.Errorf("broker is required")
	}
	kaVersion, err := sarama.ParseKafkaVersion(defaultVersion(req.Version))
	if err != nil {
		return kafkaPro, fmt.Errorf("parter version kafka error %v", err)
	}
	config := sarama.NewConfig()
	config.Version = kaVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	inProducer, err := sarama.NewAsyncProducer(req.Broker, config)
	if err != nil {
		return kafkaPro, fmt.Errorf("new async producer error %v", err)
	}
	go func() {
		for {
			select {
			case err := <-inProducer.Errors():
				zap.S().Errorf("kafka sending err", err.Msg.Key)
				inProducer.Input() <- err.Msg
			case <-inProducer.Successes():
				zap.S().Info("send queue success")
			}
		}
	}()
	kafkaPro.instanceProducer = inProducer
	kafkaPro.topics = req.Topics
	kafkaPro.ready = true
	return kafkaPro, nil
}

func initSyncKafkaProducer(req KafkaRequest) (*syncKafkaProducer, error) {
	return &syncKafkaProducer{}, nil
}

func (h *asyncKafkaProducer) Send(ctx context.Context, req KafkaPropertySend, property LoggingDTO) error {
	buffer, err := json.Marshal(property)
	if err != nil {
		return errors.New("can't marshal object")
	}
	message := &sarama.ProducerMessage{
		Topic: req.Topic,
		Key:   sarama.StringEncoder(req.Key),
		Value: sarama.ByteEncoder(buffer),
	}
	fmt.Println("send message to queue,", req.Key, req.Topic)
	h.instanceProducer.Input() <- message
	return nil
}

func (h *syncKafkaProducer) Send(ctx context.Context, req KafkaPropertySend, property LoggingDTO) error {
	return nil
}
