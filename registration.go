// File: registration.go
// Package: commandler
// Description: This file contains the registration functions for interaction commands.

package commandler

import (
	"github.com/CreativeUnicorns/dgo-commandler/utils"
	"github.com/bwmarrin/discordgo"
)

var (
	defaultDMPermission bool = false
)

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

// Register commands with Discord API (RegisterInteractionCommands)
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

// Register both handlers and global commands
func AddAndRegisterInteractionCommands(dg *discordgo.Session) {
	AddInteractionCommandHandlers(dg)
	RegisterInteractionCommands(dg)
}
