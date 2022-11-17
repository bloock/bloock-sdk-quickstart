package utils

import (
	"os"

	"github.com/fatih/color"
)

func Sample(name string, fn func(Config) error) {
	config := Config{
		ApiKey:  os.Getenv("API_KEY"),
		ApiHost: os.Getenv("API_HOST"),
	}

	color.Yellow("[+] %s: Started", name)
	err := fn(config)
	if err != nil {
		color.Red("[!] %s: Failure", name)
		color.Red("[!] %s: %s", name, err.Error())
	} else {
		color.Green("[✓] %s: Successful", name)
	}
}
