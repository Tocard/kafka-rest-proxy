package data

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func NewProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.WithField("error", err.Error()).Fatal("kafka producer failed to start")
	}

	return producer
}
