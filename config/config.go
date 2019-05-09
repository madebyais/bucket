package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Config represents config package
type Config struct {
	Local Local `yaml:"local"`
}

// Local represents configuration for local bucket
type Local struct {
	Folder string   `yaml:"folder"`
	Bucket []Bucket `yaml:"bucket"`
}

// Bucket represents bucket in local storage
type Bucket struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

// New initiates new config
func New(configFilePath string) (*Config, error) {
	configBucket := &Config{}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return configBucket, err
	}

	err = yaml.Unmarshal(configFile, configBucket)
	if err != nil {
		return configBucket, err
	}

	return configBucket, nil
}

// InitLocalBucketFolder create new local bucket folder if not exist
func InitLocalBucketFolder(folderPath string) error {
	_, err := os.Stat(folderPath)
	fmt.Println(err)
	if !os.IsNotExist(err) {
		return err
	}

	return os.MkdirAll(folderPath, 0755)
}
