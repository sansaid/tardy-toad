package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sansaid/tardy-toad/config"
)

func main() {
	appConfig := config.NewConfig()

	bot, err := discordgo.New("Bot " + appConfig.DiscordBotToken)
	// Block forever.
	select {}
}
