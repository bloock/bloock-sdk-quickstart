package utils

import (
	"os"

	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func Sample(name string, fn func(Config) error) {
	config := Config{
		ApiKey: os.Getenv("API_KEY"),
	}

	logger.Info(name + ": Started")
	err := fn(config)
	if err != nil {
		logger.Error(name + ": Failure")
		logger.Error(name + ": " + err.Error())
	} else {
		logger.Success(name + ": Successful")
	}
}
