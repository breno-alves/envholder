package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	AccessToken string `json:"access_token"`
	Project     string `json:"project,omitempty"`
}

var configPath = "/Users/brenoalves/.config/envholder/"

func ReadConfig() (*Config, error) {
	dat, err := ioutil.ReadFile(configPath + "config.json")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(dat, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func WriteConfig(accessToken string) error {
	config := &Config{
		AccessToken: accessToken,
	}
	configBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(configPath+"config.json", configBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
