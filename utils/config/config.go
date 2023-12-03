package config

import (
	"log"
	"os"
	"os/user"

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
	JwtDefaultExp int                      `yaml:"jwt_default_expiration"`
	Database      DatabaseSubconfiguration `yaml:"database"`
}

func BindConfiguration(cfg *TiseaConfiguration) *TiseaConfiguration {
	usr, _ := user.Current()
	file, err := os.ReadFile(usr.HomeDir + "/.tisea/config.yml")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = yaml.Unmarshal(file, cfg)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return cfg
}

func GetConfiguration() TiseaConfiguration {
	var cfg TiseaConfiguration
	BindConfiguration(&cfg)
	return cfg
}
