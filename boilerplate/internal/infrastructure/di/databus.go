package di

import (
	"async-arch/boilerplate/internal/infrastructure/contract"
	"async-arch/boilerplate/internal/infrastructure/di/databus"
	"async-arch/boilerplate/internal/infrastructure/di/env"
)

func NewDatabus(env env.Databus, log contract.Log) *databus.Databus {
	return databus.NewDatabus(env, log)
}

func NewConsumer(dbus *databus.Databus, topic string, groupID string, handler databus.ConsumerHandler) *databus.Consumer {
	return databus.NewConsumer(dbus, topic, groupID, handler)
}

func NewProducer(dbus *databus.Databus, topic string) *databus.Producer {
	return databus.NewProducer(dbus, topic)
}
