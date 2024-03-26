package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string     `yaml:"env" env-required:"true"`
	GRPC GRPCConfig `yaml:"grpc" env-required:"true"`
}

type GRPCConfig struct {
	Port                   int           `yaml:"port" env-required:"true"`
	Timout                 time.Duration `yaml:"timeout" env-required:"true"`
	StoragePath            string        `yaml:"storage_path" env-required:"true"`
	MaxDownloadUploadCalls int           `yaml:"max_download_upload_calls " env-default:"100"`
	MaxLOFCalls            int           `yaml:"max_lof_calls" env-default:"100"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()

	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
