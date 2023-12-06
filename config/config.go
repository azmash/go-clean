package config

import (
	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	URL HostConfig `gcfg:"URL"`
	DB  DBConfig   `gcfg:"DB"`
}

type HostConfig struct {
	Host string `gcfg:"Host"`
}

type DBConfig struct {
	Host       string `gcfg:"Host"`
	Credential string `gcfg:"Credential"`
	Port       string `gcfg:"Port"`
}

func ReadModuleConfig(cfg interface{}, path string) bool {
	ext := "config.ini"

	fname := path + ext
	err := gcfg.ReadFileInto(cfg, fname)
	if err == nil {
		return true
	}
	return false
}
