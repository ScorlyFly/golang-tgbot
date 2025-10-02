package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	BotToken      string
	AdminID       int64
	WebhookDomain string
	WebhookSecret string
	WebhookAdress string
}

var Conf *Config

func LoadConfig(config string) *Config {
	TOKEN := validateBotToken(os.Getenv("BOT_TOKEN"))
	ADMIN_ID := convertStringToInt(os.Getenv("ADMIN_ID"))
	WEBHOOK_DOMAIN := os.Getenv("WEBHOOK_DOMAIN")
	WEBHOOK_SECRET := os.Getenv("WEBHOOK_SECRET")
	WEBHOOK_ADRESS := os.Getenv("WEBHOOK_ADRESS")

	return &Config{
		BotToken: TOKEN,
		AdminID:  ADMIN_ID,

		WebhookDomain: WEBHOOK_DOMAIN,
		WebhookSecret: WEBHOOK_SECRET,
		WebhookAdress: WEBHOOK_ADRESS,
	}
}

// проверка наличие BOT_TOKEN
func validateBotToken(valid string) string {
	if valid == "" {
		log.Fatal("ERRRO: Bot token not set. ", valid)
	}
	log.Println("INFO: Bot token success loaded ...")
	return valid
}

func convertStringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
