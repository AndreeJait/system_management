package config

import (
	"strings"
	"system_management/internal/shared/constant"
)

type EnvMode string

func (e EnvMode) IsEqual(comparator string) bool {
	return strings.EqualFold(string(e), comparator)
}

const (
	Development EnvMode = "development"
	Production  EnvMode = "production"
	Staging     EnvMode = "staging"
)

func GetEnvModeByAppMode(appMode constant.AppMode) EnvMode {
	switch appMode {
	case constant.Staging:
		return Staging
	case constant.Production:
		return Production
	default:
		return Development
	}
}

type Config struct {
	Server struct {
		Host    string  `json:"host" yaml:"host"`
		Port    string  `json:"port" yaml:"port"`
		Name    string  `json:"name" yaml:"name"`
		Version string  `json:"version" yaml:"version"`
		Mode    EnvMode `json:"mode" yaml:"mode"`
	} `json:"server" yaml:"server"`

	DB struct {
		Port     string `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Host     string `json:"host" yaml:"host"`
		Password string `json:"password" yaml:"password"`
		Name     string `json:"db_name" yaml:"db_name"`
	} `json:"db" yaml:"db"`
}
