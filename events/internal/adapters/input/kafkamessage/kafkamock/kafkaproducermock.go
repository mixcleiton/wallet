package kafkamock

import "github.com/stretchr/testify/mock"

type MockKafkaProducer struct {
	mock.Mock
}

func (m *MockKafkaProducer) Producer(topic string, message string) {

}
