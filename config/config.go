package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"runtime"
	"strings"
)

func LoadConfig(appMode EnvMode) *Config {
	config := Config{}
	data, err := os.ReadFile(fmt.Sprintf(FileLocationFormat,
		strings.Split(getSourcePath(), "/config")[0], appMode))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Panicf("failed to read file %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.Panicf("failed to read file %v", err)
	}
	return &config
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
