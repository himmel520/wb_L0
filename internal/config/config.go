package config

type App struct {
	Addr string `env:"APP_ADDR" envDefault:":8080"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"postgres"`
	DBName   string `env:"POSTGRES_DB" envDefault:"order"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

type Redis struct {
	Addr string `env:"REDIS_ADDR" envDefault:"redis:6379"`
}

type Nats struct {
	Url string `env:"NATS_URL" envDefault:"nats://nats:4222"`
}

type Config struct {
	App      App
	Postgres Postgres
	Redis    Redis
	Nats     Nats
}

func NewConfig() *Config {
	return &Config{}
}
