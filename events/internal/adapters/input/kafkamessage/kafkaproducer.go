package kafkamessage

import (
	"fmt"

	"br.com.cleiton/events/internal/config"
	"github.com/IBM/sarama"
)

type kafkaProducer struct {
	kafkaConfig config.KafkaConfig
}

func NewKafkaProducer(kafkaConfig config.KafkaConfig) *kafkaProducer {
	return &kafkaProducer{kafkaConfig: kafkaConfig}
}

func (p *kafkaProducer) Producer(topic string, message string) {
	config := sarama.NewConfig()
	address := fmt.Sprintf("%s:%d", p.kafkaConfig.Host, p.kafkaConfig.Port)
	producer, err := sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

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
