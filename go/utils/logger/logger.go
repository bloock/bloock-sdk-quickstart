package logger

import "github.com/fatih/color"

func Success(message string) {
	color.Green("[✓] %s", message)
}

func Warn(message string) {
	color.Yellow("[!] %s", message)
}

func Info(message string) {
	color.Cyan("[i] %s", message)
}

func Error(message string) {
	color.Red("[x] %s", message)
}
