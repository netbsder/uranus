package utils

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

type Application struct {
	App        App                    `yaml:"app"`
	LogLevel   string                 `yaml:"log-level"`
	Info       map[string]interface{} `yaml:"info"`
	Management Management             `yaml:"management"`
	DB         DB                     `yaml:"db"`
}

type App struct {
	Name    string
	Port    uint32
	RunMode string `yaml:"run-mode"`
}

type Management struct {
	EndpointsEnabled bool     `yaml:"endpoints-enabled"`
	Security         Security `yaml:"security"`
	Endpoint         map[string]bool
}

type Security struct {
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

type DB struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

var appConfig Application
var once sync.Once

func GetAppConfig() *Application {
	once.Do(loadAppConfig)

	return &appConfig
}

func loadAppConfig() {
	// set config default
	appConfig.LogLevel = "info"
	appConfig.App.Port = 9800

	appConfig.Management.EndpointsEnabled = true
	appConfig.Management.Endpoint = make(map[string]bool)
	appConfig.Management.Endpoint["health"] = true
	appConfig.Management.Endpoint["info"] = true

	// load config
	yamlFile, err := ioutil.ReadFile("./conf/application.yaml")
	if err != nil {
		log.Fatalln("load config error: ", err)
	}

	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		log.Fatalln("parse config error!", err)
	}

	// security basic auth
	if appConfig.Management.Security.Account == "" && appConfig.Management.Security.Password == "" {
		// 随机生成
		pass := GetRandomString(6)
		appConfig.Management.Security.Account = "admin"
		appConfig.Management.Security.Password = pass
		log.Printf("management security user info => %s:%s\n", appConfig.Management.Security.Account, appConfig.Management.Security.Password)
	}

	log.Println("Load config success!", appConfig)
}
