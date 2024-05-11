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

// commandMap maps command names to their handlers for quick lookup.
var commandMap = make(map[string]*InteractionCommand)

// AddInteractionCommandHandlers adds a handler function to a discordgo.Session which will
// process incoming interaction commands by invoking the appropriate command handler.

func AddInteractionCommandHandlers(dg *discordgo.Session) {
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if cmd, exists := commandMap[i.ApplicationCommandData().Name]; exists {
				cmd.Handler(s, i)
			}
		}
	})
}

// RegisterInteractionCommands registers all InteractionCommands with the Discord API concurrently.
// It uses the properties of each InteractionCommand to create corresponding ApplicationCommands.
func RegisterInteractionCommands(dg *discordgo.Session) {
	commands := GetInteractionCommands()
	batchSize := 10 // Adjust based on empirical testing for optimal performance.
	batches := (len(commands) + batchSize - 1) / batchSize

	var wg sync.WaitGroup
	wg.Add(batches)

	for i := 0; i < len(commands); i += batchSize {
		go func(batch []*InteractionCommand) {
			defer wg.Done()
			for _, cmd := range batch {
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
				} else {
					utils.Logger.Info("Successfully registered global command", "commandName", cmd.Name)
				}
			}
		}(commands[i:min(i+batchSize, len(commands))])
	}
	wg.Wait() // Wait for all goroutines to finish
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// AddAndRegisterInteractionCommands registers both command handlers and global commands on a discordgo.Session.
// This is typically called during the initialization phase of the bot to set up all interaction commands.
func AddAndRegisterInteractionCommands(dg *discordgo.Session) {
	AddInteractionCommandHandlers(dg)
	RegisterInteractionCommands(dg)
}
