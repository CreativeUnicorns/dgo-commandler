package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CreativeUnicorns/dgo-commandler/commandler"
	_ "github.com/CreativeUnicorns/dgo-commandler/examples/ping/commands" // Import the commands package to register the commands
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Get the bot token from the environment
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Please set the DISCORD_TOKEN environment variable.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	// This is needed to initialize the command handlers
	commandler.AddAndRegisterInteractionCommands(dg)

	log.Println("Bot is now running. Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	dg.Close()
}
