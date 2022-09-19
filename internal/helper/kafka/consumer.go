package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"gopkg.in/Shopify/sarama.v1"
)

type KafkaComsumerGroup struct {
	ready         chan bool
	ConsumerGroup sarama.ConsumerGroup
	Consumer      sarama.Consumer
	MessageCh     chan *sarama.ConsumerMessage
	ErrorCh       chan *sarama.ConsumerError
}

func NewKafkaComsumerGroup(ctx context.Context, req KafkaRequest) (*KafkaComsumerGroup, error) {
	kaVersion, err := sarama.ParseKafkaVersion(defaultVersion(req.Version))
	if err != nil {
		zap.S().Errorw("Error while parsing kafka version", zap.Error(err))
	}
	config := sarama.NewConfig()
	config.Version = kaVersion
	config.Consumer.Return.Errors = true
	config.Consumer.MaxWaitTime = 3 * time.Second
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	kaConsumer := &KafkaComsumerGroup{
		ready:     make(chan bool),
		MessageCh: make(chan *sarama.ConsumerMessage),
		ErrorCh:   make(chan *sarama.ConsumerError),
	}
	singleconsumer, err := sarama.NewConsumer(req.Broker, config)
	if err != nil {
		zap.S().Errorf("Error while new consumer", zap.Error(err))
		return kaConsumer, err
	}
	kaConsumer.Consumer = singleconsumer
	groupconsumer, err := sarama.NewConsumerGroup(req.Broker, req.Group, config)
	if err != nil {
		zap.S().Errorf("Error while new consumer", zap.Error(err))
		return kaConsumer, err
	}
	kaConsumer.ConsumerGroup = groupconsumer
	go func() {
		for {
			err := groupconsumer.Consume(ctx, req.Topics, kaConsumer)
			if err != nil {
				if err == sarama.ErrClosedConsumerGroup {
					break
				}
				time.Sleep(2 * time.Second)
			}
			if ctx.Err() != nil {
				return
			}
			kaConsumer.ready = make(chan bool)
		}
	}()
	<-kaConsumer.ready

	return kaConsumer, err
}

func (kfg *KafkaComsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	close(kfg.ready)
	return nil
}

func (kaf *KafkaComsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (kaf *KafkaComsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		zap.S().Debugf("claim message", message.Key)
		session.MarkMessage(message, "")
		kaf.MessageCh <- message
	}
	return nil
}

func (k *KafkaComsumerGroup) CloseConsumerGroup() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func(sign chan os.Signal) {
		<-sign
		if k.ErrorCh != nil {
			close(k.ErrorCh)
		}
		if k.MessageCh != nil {
			close(k.MessageCh)
		}
		zap.S().Info("*** STOP KAFKA CONSUMER ***")
		if errG := k.ConsumerGroup.Close(); errG != nil {
			zap.S().Errorf("close consumer group err %v", errG)
		}

	}(sign)
}
