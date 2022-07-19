package main

import (
	"flag"
	"log"
	tgClient "telegram-bot-go/clients/telegram"
	event_consumer "telegram-bot-go/consumer/event-consumer"
	"telegram-bot-go/events/telegram"
	"telegram-bot-go/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files-storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service is up")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service crushed", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is missing")
	}

	return *token
}
