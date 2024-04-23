package cmd

import (
	"log"

	"github.com/CreativeUnicorns/dgo-commandler/pkg"
	"github.com/bwmarrin/discordgo"
)

// CommandRegistrar handles the registration of commands with the Discord API.
type CommandRegistrar struct {
	Session *discordgo.Session
}

// RegisterCommands registers all commands stored in the registry.
func (r *CommandRegistrar) RegisterCommands() {
	for _, cmd := range pkg.GetInteractionCommands() {
		appCmd := createAppCommand(cmd)
		_, err := r.Session.ApplicationCommandCreate(r.Session.State.User.ID, "", appCmd)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v\n", cmd.Name, err)
			continue
		}
		log.Printf("Successfully registered global command: %v\n", cmd.Name)
	}
}

// createAppCommand prepares the ApplicationCommand structure from an InteractionCommand.
func createAppCommand(cmd *pkg.InteractionCommand) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     cmd.Name,
		Description:              cmd.Description,
		Options:                  cmd.Options,
		DefaultMemberPermissions: &cmd.DefaultMemberPermissions,
		DMPermission:             &cmd.DMPermission,
		NSFW:                     &cmd.NSFW,
	}
}
