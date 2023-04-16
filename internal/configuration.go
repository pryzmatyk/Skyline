package internal

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	SERVER_HOST_PORT string    `default:"localhost:8080"`
	CLIENT_HOST_PORT string    `default:"ws://localhost:8080"`
	LOG_LEVEL        log.Level `default:"info log.SetLevel()"`
}

func GetConfig() (Specification, error) {
	var s Specification
	err := envconfig.Process("skyline", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s, nil
}