package kafkamessage

import (
	"encoding/json"
	"fmt"
	"time"

	"br.com.cleiton/events/internal/adapters/input/request"
	"br.com.cleiton/events/internal/config"
	"br.com.cleiton/events/internal/domain/usecases"
	"github.com/IBM/sarama"
)

type kafkaConsumer struct {
	kafkaConfig  config.KafkaConfig
	processEvent usecases.ProcessEvent
}

func NewKafkaConsumer(kafkaConfig config.KafkaConfig, processEvent usecases.ProcessEvent) *kafkaConsumer {
	return &kafkaConsumer{kafkaConfig: kafkaConfig, processEvent: processEvent}
}

func (k *kafkaConsumer) LoadReadMessages() {
	config := sarama.NewConfig()
	addressKafka := fmt.Sprintf("%s:%d", k.kafkaConfig.Host, k.kafkaConfig.Port)
	consumer, err := sarama.NewConsumer([]string{addressKafka}, config)

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Criar um consumidor para o tópico
	partitionConsumer, err := consumer.ConsumePartition("event-process", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	// Criar uma goroutine para consumir as mensagens
	go func() {
		for msg := range partitionConsumer.Messages() {
			fmt.Printf("Mensagem recebida: %s\n", string(msg.Value))
			var eventRequest request.EventRequest
			err := json.Unmarshal(msg.Value, &eventRequest)
			if err != nil {
				// Tratar erro
				fmt.Println("Erro ao deserializar JSON:", err)
				continue
			}
			k.processEvent.ProcessEvent(eventRequest.Id)
			time.Sleep(time.Second) // Simular um processamento
		}
	}()
}
