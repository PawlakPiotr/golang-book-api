package config

import (
	"encoding/json"
	"errors"
	"os"

	m "golang-book-api/model"
)

// LoadConfiguration loads confirguration from json file in ./config direcotry
func LoadConfiguration(file string) (*m.Config, error) {
	var config m.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil, errors.New("Error reading config file")
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config, nil
}
