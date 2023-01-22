package main

import (
	"andrusha/bot"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	godotenv.Load("config.env")

	var config bot.Configuration
	err := envconfig.Process("BOT", &config)
	if err != nil {
		panic(err)
	}

	bot, err := bot.New(config)
	if err != nil {
		panic(err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChannel

	defer bot.Close()
}
