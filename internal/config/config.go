package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configfilename = ".gatorconfig.json"


type Config struct {
    DbURL string `json:"db_url"`
		Username string `json:"current_user_name"`
}


func Read() (*Config, error) {
	
	configFilePath, err := getConfigFilePath()

	if err != nil {
		return nil, fmt.Errorf("unable to find config file: %v", err)
	}

	file, err := os.Open(configFilePath)

	if err != nil {
		return nil, fmt.Errorf("could not open .gatorconfig: %v", err)
	}

	defer file.Close()

	var config Config

	decoder := json.NewDecoder(file)
	
	err = decoder.Decode(&config)

	if err != nil {
		return nil, fmt.Errorf("could not decode JSON: %v", err)
	}

	return &config, nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %v", err)
	}

	fullPath := filepath.Join(homedir, configfilename)

	return fullPath, nil
}

func write(cfg *Config)  error {

	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil

}

func (cfg *Config) SetUser(username string) error{

		if len(username) == 0 {
			return fmt.Errorf("unable to find username")
		}

		cfg.Username = username

		err := write(cfg)

		if err != nil {
			return fmt.Errorf("unable to write username to config file: %v", err)
		}

		return nil
}

