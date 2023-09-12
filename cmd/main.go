package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlecSmith96/dnd-scheduler-bot/internal/adapters"
	"github.com/AlecSmith96/dnd-scheduler-bot/internal/usecases"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := adapters.NewConfig()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	discord, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatalf("unable to create discord session: %v", err)
	}

	discord.AddHandler(usecases.MessageCreate)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatalf("failed to open websocket: %v", err)
	}


	// Wait here until CTRL-C or other term signal is received.
    log.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    discord.Close()
}