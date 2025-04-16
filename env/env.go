package env

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Port         string        `env:"PORT" default:"8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"10s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"10s"`
	IntParseSize int           `env:"INT_PARSE_SIZE" default:"10"`
	IntParseBase int           `env:"INT_PARSE_BASE" default:"10"`
	DBUrl        string        `env:"DB_URL" default:"'host=localhost port=5432 user=postgres password=0901 dbname=postgres sslmode=disable"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
