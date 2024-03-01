package env

type Database struct {
	DBPort     int    `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
}

type Config struct {
	DB Database
}
