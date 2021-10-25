package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Storage StorageConf `json:"storage"`
	Logger  LoggerConf  `json:"logger"`
	HTTP    HTTPConf    `json:"http"`
}

type StorageConf struct {
	DriverName string `json:"driver_name"`
	Dsn        string `json:"dsn"`
}

type LoggerConf struct {
	Level string `json:"level"`
}

type HTTPConf struct {
	Port string `json:"port"`
}

func NewConfig(filePath string) (Config, error) {
	jsonString, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("config create:%w", err)
	}

	conf := Config{}
	if err = json.Unmarshal(jsonString, &conf); err != nil {
		return Config{}, fmt.Errorf("config create:%w", err)
	}

	return conf, nil
}
