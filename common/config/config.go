package config

import "github.com/go-ini/ini"

type Application struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	Description    string `json:"description"`
	EndpointPrefix string `json:"endpoint_prefix"`
	IsInTesting    bool   `json:"is_testing"`
	Port           string `json:"port"`
	JwtSecret      string `json:"jwt_secret"`
}

type Database struct {
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Pass        string `json:"pass"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

var config *ini.File
var DatabaseConfig = &Database{}
var ApplicationConfig = &Application{}

func init() {
	var err error
	config, err = ini.Load("config.ini")
	if err != nil {
		panic(err)
	}

	mapTo("Application", ApplicationConfig)
	mapTo("Database", DatabaseConfig)

}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}
