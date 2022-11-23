package main

import (
	"TelegramForArbitrage/api/server"
	"TelegramForArbitrage/pkg/db"
	"TelegramForArbitrage/pkg/telegram"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"os"
)

func main() {
	log.Init("telegramApi", true)
	telegramApi := new(telegram.Telegram)
	var err error

	telegramApiToken := os.Getenv("TELEGRAMAPITOKEN")

	telegramApi.BotApi, err = telegram.ConnectToTelegram(telegramApiToken)
	if err != nil {
		log.Error.Fatal(err)
	}

	telegramApi.MongoConnector, err = db.ConnectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Error.Fatal(err)
	}

	go telegramApi.CheckUsersStatus()

	if err = server.StartGrpcServer(telegramApi); err != nil {
		log.Error.Fatal(err)
	}

}
