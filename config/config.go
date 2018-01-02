package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	MaxConnections int                 `yaml:"MaxConnections"`
	MaxMessageSize int                 `yaml:"MaxMessageSize"`
	Timeout        float64             `yaml:"Timeout"`
	UDPListeners   []UDPListenerConfig `yaml:"UDPListeners"`
}

type UDPListenerConfig struct {
	Ip   string
	Port string
}

var configFileName = "./server_config.yaml"

func CheckError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func Read() ServerConfig {

	isFile, err := exists(configFileName)
	CheckError(err)

	if !isFile {
		defaultData, err := yaml.Marshal(DefaultServerConfig())
		CheckError(err)
		err = writeConfig(defaultData)
		CheckError(err)
	}

	b, err := readConfig()
	CheckError(err)

	config := ServerConfig{}

	err = yaml.Unmarshal(b, &config)
	CheckError(err)

	return config
}

func writeConfig(data []byte) error {
	return ioutil.WriteFile(configFileName, data, 0644)
}

func readConfig() ([]byte, error) {
	return ioutil.ReadFile(configFileName)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func DefaultServerConfig() ServerConfig {
	config := ServerConfig{}
	config.MaxConnections = 5
	config.MaxMessageSize = 512000
	config.Timeout = 5.0
	config.UDPListeners = make([]UDPListenerConfig, 0)
	config.UDPListeners = append(config.UDPListeners, UDPListenerConfig{"", "10001"})
	return config
}
