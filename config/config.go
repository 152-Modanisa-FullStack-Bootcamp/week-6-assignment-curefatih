package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ServerAddr           string  `json:"serverAddr"`
	InitialBalanceAmount float32 `json:"initialBalanceAmount"`
	MinimumBalanceAmount float32 `json:"minimumBalanceAmount"`
}

var config = &Config{}

func init() {
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, config)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return config
}
