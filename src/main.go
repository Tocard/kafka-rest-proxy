package main

import (
	"data"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"server"
)

func main() {

	config := data.NewConfig()
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("config.yml file missing", err.Error())

	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error during parsing config.yml", err.Error())
	}
	log.Info(config)
	server.GetServer().Run()
}
