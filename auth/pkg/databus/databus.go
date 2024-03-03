package databus

import (
	"fmt"

	"async-arch/auth/internal/infrastructure/contract"
	"async-arch/auth/pkg/env"
)

type Databus struct {
	broker string
	log    contract.Log
}

func NewDatabus(conf env.Databus, log contract.Log) *Databus {
	return &Databus{
		broker: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		log:    log,
	}
}
