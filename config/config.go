package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

type Configuration struct {
	DatabaseType string
	DatabaseURL  string
	DatabaseName string
}

var globalConfig *Configuration

func GetDatabaseType() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.DatabaseType
}

func GetDatabaseURL() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.DatabaseURL
}

func GetDatabaseName() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.DatabaseName
}

func loadConfiguration() *Configuration {
	env := os.Getenv("APP_ENV")
	var filepath string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	if env == "DEV" || env == "" {
		filepath = path.Dir(filename) + "/dev.json"
	}

	if env == "TEST" {
		filepath = path.Dir(filename) + "/unittest.json"
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	return &config
}
