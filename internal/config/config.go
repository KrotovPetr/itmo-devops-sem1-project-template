package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

func Load(path, format string) (Config, error) {
	log.Printf("Loading configuration from '%s'\n", path)

	if err := setupViper(path, format); err != nil {
		return Config{}, fmt.Errorf("failed to setup viper: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	log.Println("Configuration loaded successfully")
	return config, nil
}

func setupViper(path, format string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType(format)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
