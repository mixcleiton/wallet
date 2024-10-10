package kafkamessage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"br.com.cleiton/events/internal/adapters/input/requestdto"
	"br.com.cleiton/events/internal/config"
	"br.com.cleiton/events/internal/domain/entities"
	"br.com.cleiton/events/internal/domain/usecases"
	"github.com/Shopify/sarama"
)

type kafkaConsumer struct {
	kafkaConfig  config.KafkaConfig
	processEvent usecases.ProcessEvent
}

func NewKafkaConsumer(kafkaConfig config.KafkaConfig, processEvent usecases.ProcessEvent) *kafkaConsumer {
	return &kafkaConsumer{kafkaConfig: kafkaConfig, processEvent: processEvent}
}

func (k *kafkaConsumer) LoadReadMessages() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	config := sarama.NewConfig()
	//config.Version = sarama.V0_11_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	addressKafka := fmt.Sprintf("%s:%d", k.kafkaConfig.Host, k.kafkaConfig.Port)
	log.Println("address kafka" + addressKafka)
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)

	if err != nil {
		panic(err)
	}

	// Criar um consumidor para o tópico
	partitionConsumer, err := consumer.ConsumePartition("event-process", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	// Criar uma goroutine para consumir as mensagens
	go func() {
		for {
			log.Println("processando")
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("Mensagem recebida: %s\n", string(msg.Value))
				var eventRequest requestdto.EventRequest
				err := json.Unmarshal(msg.Value, &eventRequest)
				if err != nil {
					// Tratar erro
					fmt.Println("Erro ao deserializar JSON:", err)
					continue
				}
				log.Println("eventRequestId ", eventRequest.UUID)
				event := entities.Event{
					Id:     eventRequest.Id,
					IdUUID: eventRequest.UUID,
				}
				k.processEvent.ProcessEvent(&event)
			}

			time.Sleep(time.Second) // Simular um processamento
		}
	}()
}
