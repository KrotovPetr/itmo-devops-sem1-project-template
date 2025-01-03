package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path, format string) (Config, error) {
    log.Printf("Loading configuration from '%s'\n", path)

    viper.SetConfigFile(path)
    viper.SetConfigType(format)

    if err := viper.ReadInConfig(); err != nil {
        return Config{}, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return Config{}, err
    }

    log.Println("Configuration loaded successfully")
    
    return config, nil
}