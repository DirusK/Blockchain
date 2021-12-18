package config

import (
	"log"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v3"
)

//go:generate go-validator

const (
	// DefaultPath - default path for config.
	DefaultPath = "../cmd/config.yaml"
)

type (
	// Config defines the properties of the application configuration.
	Config struct {
		HTTPServer HTTPServer `yaml:"http-server" valid:"required"`
	}

	// HTTPServer defines HTTP section of the API server configuration.
	HTTPServer struct {
		ListenAddress   string        `yaml:"listen-address" valid:"required"`
		ReadTimeout     time.Duration `yaml:"read-timeout" valid:"required"`
		WriteTimeout    time.Duration `yaml:"write-timeout" valid:"required"`
		HeaderLimit     int           `yaml:"header-limit" valid:"required"`
		GracefulTimeout int           `yaml:"graceful-timeout" valid:"required"`
	}
)

// New loads and validates all configuration data, returns filled Cfg - configuration data model.
func New(cfgFilePath string) (*Config, error) {
	cfg := &Config{}

	data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	_, err = govalidator.ValidateStruct(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return cfg, err
}
