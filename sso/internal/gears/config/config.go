package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const CfgEnvVar = "KONIO_SSO_CFG"
const CfgFlag = "config"

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	GRPC    GRPCConfig    `yaml:"grpc"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	Env      string        `yaml:"env" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type StorageConfig struct {
	Name string `yaml:"name" env-required:"true"`
}

func MustLoad() *Config {
	cfgPath, err := configPath()
	if err != nil {
		panic(fmt.Errorf("error: failed to load config: %w", err).Error())
	}

	if _, err = os.Stat(cfgPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("error: config file %s does not exist", cfgPath))
	}

	cfg := new(Config)

	if err = cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		panic("error: failed to read config" + err.Error())
	}

	return cfg
}

func configPath() (string, error) {
	var cfgPath string

	cfgPath = os.Getenv(CfgEnvVar)
	if strings.TrimSpace(cfgPath) != "" {
		return cfgPath, nil
	}

	flag.StringVar(&cfgPath, CfgFlag, "", "set path to configuration file")
	flag.Parse()

	if strings.TrimSpace(cfgPath) != "" {
		return cfgPath, nil
	}

	return "", fmt.Errorf(
		"config path is not defined.\n set env variable %s\n or set flag '%s'",
		CfgEnvVar,
		CfgFlag,
	)
}
