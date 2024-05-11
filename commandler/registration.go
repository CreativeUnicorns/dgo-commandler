// Package commandler provides structures and functions to manage the registration and execution
// of Discord interaction commands.
package commandler

import (
	"sync"

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
		// Ensure the interaction is a command before processing
		if i.Type == discordgo.InteractionApplicationCommand {
			for _, cmd := range GetInteractionCommands() {
				utils.Logger.Info("Registering handler for command", "commandName", cmd.Name)
				if i.ApplicationCommandData().Name == cmd.Name {
					cmd.Handler(s, i)
					break
				}
			}
		}
	})
}

// RegisterInteractionCommands registers all InteractionCommands with the Discord API concurrently.
// It uses the properties of each InteractionCommand to create corresponding ApplicationCommands.
func RegisterInteractionCommands(dg *discordgo.Session) {
	commands := GetInteractionCommands()
	var wg sync.WaitGroup
	wg.Add(len(commands))

	for _, cmd := range commands {
		go func(cmd InteractionCommand) {
			defer wg.Done()
			appCmd := &discordgo.ApplicationCommand{
				Name:         cmd.Name,
				Description:  cmd.Description,
				DMPermission: &defaultDMPermission,
				Options:      cmd.Options,
			}

			if cmd.DefaultMemberPermissions != 0 {
				appCmd.DefaultMemberPermissions = &cmd.DefaultMemberPermissions
			}

			if cmd.DMPermission {
				appCmd.DMPermission = &cmd.DMPermission
			}

			if cmd.NSFW {
				appCmd.NSFW = &cmd.NSFW
			}

			_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", appCmd)
			if err != nil {
				utils.Logger.Error("Cannot create command", "commandName", cmd.Name, "error", err)
				return
			}
			utils.Logger.Info("Successfully registered global command", "commandName", cmd.Name)
		}(*cmd)
	}
	wg.Wait() // Wait for all goroutines to finish
}

// AddAndRegisterInteractionCommands registers both command handlers and global commands on a discordgo.Session.
// This is typically called during the initialization phase of the bot to set up all interaction commands.
func AddAndRegisterInteractionCommands(dg *discordgo.Session) {
	AddInteractionCommandHandlers(dg)
	RegisterInteractionCommands(dg)
}
