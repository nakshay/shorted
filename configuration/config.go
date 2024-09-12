package configuration

import (
	"encoding/json"
	"os"
)

type ConfigData struct {
	ServiceDomain         string `json:"service_domain"`
	RandomCharacterLength int    `json:"random_char_length"`
}

type ConfigLoader interface {
	LoadConfig(file string) (*ConfigData, error)
}

type configLoader struct{}

func NewConfigLoader() ConfigLoader {
	return configLoader{}
}

func (c configLoader) LoadConfig(filePath string) (*ConfigData, error) {
	configuration := ConfigData{}
	file, err := os.Open(filePath)
	if err != nil {
		return &configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return &configuration, err
}
