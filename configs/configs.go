package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	TTLValue string `env:"LOCATION_HISTORY_TTL_SECONDS"`
	Port     string `env:"HISTORY_SERVER_LISTEN_ADDR" envDefault:":8080"`
}

func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.TTLValue),
		v.Field(&c.Port),
	)
}

func InitConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg
}
