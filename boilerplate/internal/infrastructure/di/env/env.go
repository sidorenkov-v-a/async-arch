package env

type Database struct {
	DBUrl string `env:"DB_URL"`
}

type Config struct {
	DB Database
}
