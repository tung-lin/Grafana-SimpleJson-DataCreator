package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Port        string `yaml:port`
	DB_Address  string `yaml:db_address`
	DB_User     string `yaml:db_user`
	DB_Password string `yaml:db_password`
}

var Config Configuration

func init() {

	configFilePath, _ := os.Getwd()
	file, err := ioutil.ReadFile(configFilePath + "/datacreator.yaml")

	if err != nil {
		log.Fatalf("Read config file failed: %v", err)
	}

	err = yaml.Unmarshal(file, &Config)

	if err != nil {
		log.Fatalf("Unmarshal config file failed: %v", err)
	}

	GetEnv("port", &Config.Port)
	GetEnv("db_address", &Config.DB_Address)
	GetEnv("db_user", &Config.DB_User)
	GetEnv("db_password", &Config.DB_Password)
}

func GetEnv(envName string, variable *string) {

	envName = strings.ToLower(envName)
	value, existed := os.LookupEnv(envName)

	if existed {
		*variable = value
	}
}
