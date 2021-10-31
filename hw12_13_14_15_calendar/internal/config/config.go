package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	config         Config
	once           sync.Once
	configFilePath = "./configs/default.json"
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

func SetFilePath(filePath string) {
	configFilePath = filePath
}

func Get() *Config {
	once.Do(func() {
		jsonString, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			log.Fatal(err)
		}

		if err = json.Unmarshal(jsonString, &config); err != nil {
			log.Fatal(err)
		}
	})
	return &config
}
