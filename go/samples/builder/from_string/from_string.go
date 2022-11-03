package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-quickstart/utils"
)

func main() {
	utils.Sample("builder_from_string", func(c utils.Config) error {
		return errors.New("an error")
	})
}
