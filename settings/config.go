package settings

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Database struct {
		File     string `yaml:"file" envconfig:"DB_FILE"`
		Username string `yaml:"sql-user" envconfig:"DB_USERNAME"`
		Password string `yaml:"sql-pass" envconfig:"DB_PASSWORD"`
		Database string `yaml:"sql-database" envconfig:"DB_DATABASE"`
		Server   string `yaml:"sql-server" envconfig:"DB_SERVER"`
		Port     int    `yaml:"sql-port" envconfig:"DB_PORT"`
	} `yaml:"database"`

	Redis struct {
		Server string `yaml:"server" envconfig:"REDIS_SERVER"`
	} `yaml:"redis"`

	SSO struct {
		ClientId        string `yaml:"clientId" envconfig:"SSO_CLIENTID"`
		ClientSecret    string `yaml:"clientSecret" envconfig:"SSO_CLIENTSECRET"`
		BaseAddress     string `yaml:"baseAddress" envconfig:"SSO_BASEADDRESS"`
		BackBaseAddress string `yaml:"backBaseAddress" envconfig:"SSO_BACKBASEADDRESS"`
	} `yaml:"sso"`
}

func ReadConfig(cfg *Config) {
	readFile(cfg)
	readEnv(cfg)
	fmt.Printf("%+v", cfg)
}

func readFile(cfg *Config) {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get working directory: %v", err))
	}

	// Construct the absolute path to config.yml
	fileName := filepath.Join(dir, "config.yml")

	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failed to open config file: %s, error: %v", fileName, err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close config file: %s, error: %v", fileName, err))
		}
	}(f)

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		panic(fmt.Sprintf("Failed to decode config file: %s, error: %v", fileName, err))
	}

	fmt.Printf("Loaded Config: %+v\n", cfg)
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to process environment variables: %v", err))
	}
}
