package di

import (
	"async-arch/task_tracker/internal/infrastructure/contract"
	databusClient "async-arch/task_tracker/pkg/databus"
	"async-arch/task_tracker/pkg/env"
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
