package config

import (
	"time"
)

const CfgEnvVar = "KONIO_SSO_CFG"
const CfgFlag = "config"

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
