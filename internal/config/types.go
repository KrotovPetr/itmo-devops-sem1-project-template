package config

import (
	"time"
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
