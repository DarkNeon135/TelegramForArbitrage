package main

import (
	"TelegramMessagesSender/pkg/db"
	telegram2 "TelegramMessagesSender/pkg/telegram"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"os"
)

func main() {
	log.Init("telegram", true)
	telegram := new(telegram2.Telegram)
	var err error

	telegramApiToken := os.Getenv("TELEGRAMAPITOKEN")

	telegram.BotApi, err = telegram2.ConnectToTelegram(telegramApiToken)
	if err != nil {
		log.Error.Fatal(err)
	}

	telegram.MongoConnector, err = db.ConnectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Error.Fatal(err)
	}

	telegram.CheckUsersStatus()

}
