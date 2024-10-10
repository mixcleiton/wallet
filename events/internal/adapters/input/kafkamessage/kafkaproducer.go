package kafkamessage

import (
	"fmt"
	"log"
	"os"

	"br.com.cleiton/events/internal/config"
	"github.com/Shopify/sarama"
)

type kafkaProducer struct {
	kafkaConfig config.KafkaConfig
}

func NewKafkaProducer(kafkaConfig config.KafkaConfig) *kafkaProducer {
	return &kafkaProducer{kafkaConfig: kafkaConfig}
}

func (p *kafkaProducer) Producer(topic string, message string) {

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada para a partição %d a offset %d\n", partition, offset)
}
