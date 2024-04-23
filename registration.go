// File: registration.go
// Package: commandler
// Description: This file contains the registration functions for interaction commands.

package commandler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	defaultDMPermission bool = false
)

func AddInteractionCommandHandlers(dg *discordgo.Session) {
	// Register interaction handler
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		for _, cmd := range GetInteractionCommands() {
			fmt.Printf("Registering handler for command: %v\n", cmd.Name)
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
			log.Printf("Cannot create '%v' command: %v\n", cmd.Name, err)
			continue
		}
		fmt.Printf("Successfully registered global command: %v\n", cmd.Name)
	}
}

// Register both handlers and global commands
func AddAndRegisterInteractionCommands(dg *discordgo.Session) {
	AddInteractionCommandHandlers(dg)
	RegisterInteractionCommands(dg)
}
