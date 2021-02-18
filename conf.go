package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Структура конфигурационного файла
type Config struct {
	Imap struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Server   string `yaml:"server"`
	}
	LogLevel string `yaml:"loglevel"`
}

func (a *MyApp) GetConfigYaml(filename string) {
	log.Info("Reading config ", filename)

	var conf Config

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	switch conf.LogLevel {
	case "debug":
		a.logLevel = log.DebugLevel
	case "info":
		a.logLevel = log.InfoLevel
	case "warn":
		a.logLevel = log.WarnLevel
	case "error":
		a.logLevel = log.ErrorLevel
	case "fatal":
		a.logLevel = log.FatalLevel
	default:
		a.logLevel = log.InfoLevel
	}

	a.imapUsername = conf.Imap.Username
	a.imapPassword = conf.Imap.Password
	a.imapServer = conf.Imap.Server

	log.Info("LogLevel: ", conf.LogLevel)
}
