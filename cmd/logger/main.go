package main

import (
	"bot/internal/logger"
	"fmt"
)

func main() {
	logging := logger.LoadLogger()
	logging.Debug = false
	fmt.Println(logging.Debug)
	logging.Info()
}
