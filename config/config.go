package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB   Postgres `yaml:"postgres"`
	Http HTTP     `yaml:"http_server"`
	Env  string   `yaml:"env"`
	Jwt  string   `yaml:"jwt"`
}
type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `env:"POSTGRES_PASS"`
	DbName   string `yaml:"db_name"`
	SslMode  string `yaml:"ssl_mode"`
	TimeZone string `yaml:"time_zone"`
}

func (db *Postgres) GetDNS() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		db.Host, db.User, db.Password, db.DbName, db.Port, db.SslMode, db.TimeZone)
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error yaml: %w", err)
	}

	err = cleanenv.ReadConfig("./config/config.env", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error env: %w", err)
	}

	return cfg, nil
}
