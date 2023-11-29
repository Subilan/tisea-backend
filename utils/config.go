package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseSubconfiguration struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"db"`
}

type TiseaConfiguration struct {
	JwtPrivateKey string                   `yaml:"jwt_private_key"`
	JwtExpiration int                      `yaml:"jwt_expiration"`
	Database      DatabaseSubconfiguration `yaml:"database"`
}

func BindConfiguration(cfg *TiseaConfiguration) *TiseaConfiguration {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(file, cfg)

	if err != nil {
		return nil
	}

	return cfg
}

func GetConfiguration() TiseaConfiguration {
	var cfg TiseaConfiguration
	BindConfiguration(&cfg)
	return cfg
}
