package env

type Database struct {
	DBPort     int    `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
}

type Databus struct {
	Host string `env:"DATABUS_HOST"`
	Port int    `env:"DATABUS_PORT"`
}

type Server struct {
	BindAddr string `env:"SERVER_BIND_ADDR"`
}

type Config struct {
	DB      Database
	Databus Databus
	Server  Server
}
