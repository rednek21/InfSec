package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	ENV    string       `yaml:"env"`
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
}

func MustLoad() (*Config, error) {
	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	file, err := os.Open(path)
	if err != nil {
		panic("failed to open config file: " + err.Error())
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		panic("failed to parse config file: " + err.Error())
	}

	return &cfg, nil
}

// Read --config=path/to/config or try to load config from env
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG")
	}

	return res
}
