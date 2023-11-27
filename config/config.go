package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig   appConfig
	DbConfig    databaseConfig
	RedisConfig redisConfig
	KeyConfig   keyConfig
}

type appConfig struct {
	AppHost      string `env:"APP_HOST"`
	Port         string `env:"APP_PORT,required=true"`
	WriteTimeout uint32 `env:"WRITE_TIMEOUT,required=true"`
	ReadTimeout  uint32 `env:"READ_TIMEOUT,required=true"`
	IdleTimeout  uint32 `env:"IDLE_TIMEOUT,required=true"`
}
type databaseConfig struct {
	DB_Host    string `env:"DB_HOST,required=true"`
	DB_User    string `env:"DB_USER,required=true"`
	DB_Name    string `env:"DB_NAME,required=true"`
	DB_Pswd    string `env:"DB_PASSWORD,required=true"`
	DB_Port    string `env:"DB_PORT,required=true"`
	DB_Sslmode string `env:"DB_SSLMODE,required=true"`
}
type redisConfig struct {
	Addr     string `env:"Address,required=true"`
	Password string `env:"Password"`
	DB       int    `env:"DB"`
}
type keyConfig struct {
	Public_Key  string `env:"Public_Key,required=true"`
	Private_Key string `env:"Private_Key,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
