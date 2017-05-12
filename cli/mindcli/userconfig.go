package mindcli

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type UserConfig struct {
	DefaultRobotName string  `json:"DefaultRobotName"`
	Robots           []Robot `json:"Robots"`
	DockerImage      string  `json:"DockerImage"`
	path             string
	authFile         string
	userHash         string
}

func NewUserConfig(path, authFile string) *UserConfig {
	var config UserConfig
	file, err := os.Open(path)
	if err != nil {
		return &UserConfig{path: path, authFile: authFile}
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return &UserConfig{path: path, authFile: authFile}
	}
	config.path = path
	config.authFile = authFile
	userHashBytes, _ := ioutil.ReadFile(config.authFile)
	config.userHash = string(userHashBytes)
	return &config
}

func (config *UserConfig) Write() error {
	err := ioutil.WriteFile(config.authFile, []byte(config.userHash), 0600)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.path, configJson, 0644)
}
