package main

import (
	"log"
	"os"

	"github.com/CreativeUnicorns/dgo-commandler/cmd"
	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("No token provided")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}

	handler := cmd.InteractionHandler{Session: session}
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler.HandleCommand(i)
	})

	err = session.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}
