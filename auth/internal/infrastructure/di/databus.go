package di

import (
	"async-arch/boilerplate/internal/infrastructure/contract"
	databusClient "async-arch/boilerplate/pkg/databus"
	"async-arch/boilerplate/pkg/env"
)

func NewDatabus(env env.Databus, log contract.Log) *databusClient.Databus {
	return databusClient.NewDatabus(env, log)
}

func NewConsumer(dbus *databusClient.Databus, topic string, groupID string, handler databusClient.ConsumerHandler) *databusClient.Consumer {
	return databusClient.NewConsumer(dbus, topic, groupID, handler)
}

func NewProducer(dbus *databusClient.Databus, topic string) *databusClient.Producer {
	return databusClient.NewProducer(dbus, topic)
}
