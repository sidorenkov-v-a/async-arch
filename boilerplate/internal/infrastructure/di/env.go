package di

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"async-arch/boilerplate/internal/infrastructure/di/env"
)

func NewEnv() (*env.Config, error) {
	var cfg env.Config

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig(fmt.Sprintf("%s/.env", pwd), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
