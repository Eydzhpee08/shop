package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


type Config struct {
	Server Server `json:"server"`
	DSN    string `json:"db_dsn"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func InitConfigs() Config {
	
	bytes, err:=ioutil.ReadFile("./appConfig.json")
	if err != nil {
		panic(err)
	}

	config:=Config{}
	err=json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	log.Println(config)
	return config
}