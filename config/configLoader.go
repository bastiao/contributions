package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

/**
* This is the module to load configuration
 */

// Configuration for Phabricator
type ConfPha struct {
	Endpoint string `yaml:"endpoint"`
	Token    string `yaml:"token"`
	Repo     string `yaml:"repo"`
}

// Configuration for Jenkins
type ConfJenkins struct {
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
	Pipeline string `yaml:"pipeline"`
}

// Configurations for IMAP
type ConfImap struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address string `yaml:"address"`
}

// Configurations General
type ConfGoPath struct {
	PhaConf    ConfPha     `yaml:"phabricator"`
	PhaJenkins ConfJenkins `yaml:"jenkins"`
	PhaImap    ConfImap    `yaml:"imap"`
}

// Load from File
func (c *ConfGoPath) FromFile(file string) *ConfGoPath {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
