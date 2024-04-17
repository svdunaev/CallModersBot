package main

import (
	"github.com/joho/godotenv"
	"log"
	"nattiAlertBot/clients/telegram"
	"nattiAlertBot/internal/config"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()
	tgClient := telegram.New(cfg.Token, cfg.Host)
	//TODO: fetcher.New(tgClient)
	//TODO: processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}
