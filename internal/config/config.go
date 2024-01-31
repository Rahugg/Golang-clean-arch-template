package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HttpServer HttpServer `yaml:"HttpServer"`
	Database   Database   `yaml:"Database"`
	Auth       Auth       `yaml:"JwtSecretKey"`
	Transport  Transport  `yaml:"Transport"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type Database struct {
	Main    DbNode `yaml:"Main"`
}

type DbNode struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}

type Auth struct {
	PasswordSecretKey string `yaml:"PasswordSecretKey"`
	JwtSecretKey      string `yaml:"JwtSecretKey"`
}

type Transport struct {
	User     UserTransport     `yaml:"user"`
}

type UserTransport struct {
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

func New(path string) (config Config, err error) {
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	
		viper.AutomaticEnv()
	
		err = viper.ReadInConfig()
		if err != nil {
			return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
		}
	
		err = viper.Unmarshal(&config)
		if err != nil {
			return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
		}
	
		return config, nil
}
