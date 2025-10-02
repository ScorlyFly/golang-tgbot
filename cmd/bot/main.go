package main

import (
	"bot/internal/bot"
	"log"

	"github.com/joho/godotenv"
)

const (
	config = "configs/config.yaml"
	env    = "configs/.env"
)

func envload() {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("INFO: env file success loaded.")
}

func main() {
	envload() // -- убераем есили нам ненужен env
	bot.Run(config)
}
