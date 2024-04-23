// Package commandler provides structures and functions to manage the registration and execution
// of Discord interaction commands.
package commandler

import (
	"github.com/CreativeUnicorns/dgo-commandler/utils"
	"github.com/bwmarrin/discordgo"
)

// defaultDMPermission specifies the default permission for direct messages when registering commands.
var defaultDMPermission bool = false

// AddInteractionCommandHandlers adds a handler function to a discordgo.Session which will
// process incoming interaction commands by invoking the appropriate command handler.
func AddInteractionCommandHandlers(dg *discordgo.Session) {
	// Register interaction handler
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		for _, cmd := range GetInteractionCommands() {
			utils.Logger.Info("Registering handler for command", "commandName", cmd.Name)
			if i.ApplicationCommandData().Name == cmd.Name {
				cmd.Handler(s, i)
				break
			}
		}
	})
}

// RegisterInteractionCommands registers all InteractionCommands with the Discord API.
// It uses the properties of each InteractionCommand to create corresponding ApplicationCommands.
func RegisterInteractionCommands(dg *discordgo.Session) {
	for _, cmd := range GetInteractionCommands() {
		// Prepare the application command to create
		appCmd := &discordgo.ApplicationCommand{
			Name:         cmd.Name,
			Description:  cmd.Description,
			DMPermission: &defaultDMPermission,
			Options:      cmd.Options,
		}

		// Only add DefaultMemberPermissions if set (non-zero)
		if cmd.DefaultMemberPermissions != 0 {
			appCmd.DefaultMemberPermissions = &cmd.DefaultMemberPermissions
		}

		// Only add DMPermission if it differs from the default
		if cmd.DMPermission {
			appCmd.DMPermission = &cmd.DMPermission
		}

		// Only add NSFW if true since default is false
		if cmd.NSFW {
			appCmd.NSFW = &cmd.NSFW
		}

		// Create the command globally
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", appCmd)
		if err != nil {
			utils.Logger.Error("Cannot create command", "commandName", cmd.Name, "error", err)
			continue
		}
		utils.Logger.Info("Successfully registered global command", "commandName", cmd.Name)
	}
}

// AddAndRegisterInteractionCommands registers both command handlers and global commands on a discordgo.Session.
// This is typically called during the initialization phase of the bot to set up all interaction commands.
func AddAndRegisterInteractionCommands(dg *discordgo.Session) {
	AddInteractionCommandHandlers(dg)
	RegisterInteractionCommands(dg)
}
