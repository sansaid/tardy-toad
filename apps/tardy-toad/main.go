package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sansaid/tardy-toad/config"
)

var (
	commands = []discordgo.ApplicationCommand{
		{
			Name:        "tardy-toad",
			Description: "Manage tardy toads",
		},
	}
)

func main() {
	appConfig := config.NewConfig()

	bot, err := discordgo.New("Bot " + appConfig.DiscordBotToken)

	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	// TODO: use subcommands to implement - https://discordnet.dev/guides/int_basics/application-commands/slash-commands/subcommands.html
	// /tardy-toad show -- list link to tardy toads
	// /tardy-toad report -- report a tardy toad (modal interaction)
	// /tardy-toad delete -- delete a tardy toad entry (use options)

	bot.AddHandler()

	// Block forever.
	select {}
}
