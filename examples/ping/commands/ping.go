package commands

import (
	"github.com/CreativeUnicorns/dgo-commandler/commandler"
	"github.com/bwmarrin/discordgo"
)

func init() {
	pingCommand := &commandler.InteractionCommand{
		Name:        "ping",
		Description: "Responds with Pong!",
		Handler:     pingHandler,
	}
	commandler.RegisterInteractionCommand(pingCommand)
}

// pingHandler is the handler function for the ping command
func pingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
