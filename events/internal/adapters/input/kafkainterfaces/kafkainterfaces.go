package kafkainterfaces

type KafkaConsumerInterface interface {
	LoadReadMessages()
}

type KafkaProducerInterface interface {
	Producer(topic string, message string)
}
