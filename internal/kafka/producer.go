package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

const (
	CreateRuleTopic = "create_rule"
	UpdateRuleTopic = "update_rule"
	RemoveRuleTopic = "remove_rule"
)

type AsyncProducer interface {
	SendMessageWithContext(ctx context.Context, msg *sarama.ProducerMessage)
}

type asyncProducer struct {
	sarama.AsyncProducer
}

func NewAsyncProducer(brokers []string) (*asyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	p, err := sarama.NewAsyncProducer(brokers, config)

	producer := &asyncProducer{
		AsyncProducer: p,
	}

	return producer, err
}

func (p *asyncProducer) SendMessageWithContext(ctx context.Context, msg *sarama.ProducerMessage) {
	select {
	case p.AsyncProducer.Input() <- msg:
	case <-ctx.Done():
	}
}

func PrepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}
