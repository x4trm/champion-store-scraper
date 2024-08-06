package config

import (
	"encoding/json"
	"os"
)

type ConfigEntry struct {
	Item     string `json:"item"`
	BaseUrl  string `json:"baseUrl"`
	MaxPages int    `json:"maxPages"`
	Sex      string `json:"sex"`
}

type Config struct {
	Entries []ConfigEntry `json:"entries"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
